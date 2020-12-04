package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/akkagao/gms/common"
)

/*
MessagePack 消息编码、解码
实现gnet.ICodec 接口
*/
type MessagePack struct {
	// Message
}

func NewMessagePack() IMessagePack {
	return &MessagePack{}
}

/*
Pack 消息编码
消息格式
扩展数据长度|主体数据长度|扩展数据|主体数据
*/
func (m *MessagePack) Pack(message Imessage) ([]byte, error) {
	result := make([]byte, 0)

	buffer := bytes.NewBuffer(result)

	if err := binary.Write(buffer, binary.BigEndian, message.GetHeader()); err != nil {
		return nil, err
	}

	serviceFuncL := len(message.ServiceFunc())

	extData := encodeExt(message.GetExt())
	extDataL := len(extData)

	data := message.GetData()
	dataL := len(data)

	// 写入消息总长
	totalL := (4 + serviceFuncL) + + (4 + extDataL) + (4 + dataL)
	// fmt.Println("totleL:", totalL)
	if err := binary.Write(buffer, binary.BigEndian, uint32(totalL)); err != nil {
		return nil, err
	}
	// 写入方法名总长
	if err := binary.Write(buffer, binary.BigEndian, uint32(serviceFuncL)); err != nil {
		return nil, err
	}
	// 写入方法名
	if err := binary.Write(buffer, binary.BigEndian, []byte(message.ServiceFunc())); err != nil {
		return nil, err
	}

	// 写入扩展信息长度
	if err := binary.Write(buffer, binary.BigEndian, uint32(extDataL)); err != nil {
		return nil, err
	}
	// 写入扩展信息
	if err := binary.Write(buffer, binary.BigEndian, extData); err != nil {
		return nil, err
	}

	// 写入消息内容长度
	if err := binary.Write(buffer, binary.BigEndian, uint32(dataL)); err != nil {
		return nil, err
	}
	// 写入消息内容
	if err := binary.Write(buffer, binary.BigEndian, data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// len,string,len,string,......
func encodeExt(m map[string]string) []byte {
	if len(m) == 0 {
		return []byte{}
	}
	var buf bytes.Buffer
	var d = make([]byte, 4)
	for k, v := range m {
		binary.BigEndian.PutUint32(d, uint32(len(k)))
		buf.Write(d)
		buf.Write(common.StringToSliceByte(k))
		binary.BigEndian.PutUint32(d, uint32(len(v)))
		buf.Write(d)
		buf.Write(common.StringToSliceByte(v))
	}
	return buf.Bytes()
}

func decodeExt(l uint32, data []byte) (map[string]string, error) {
	m := make(map[string]string, 10)
	n := uint32(0)
	for n < l {
		// parse one key and value
		// key
		sl := binary.BigEndian.Uint32(data[n : n+4])
		n = n + 4
		if n+sl > l-4 {
			return m, errors.New("wrong ext some keys or values are missing")
		}
		k := string(data[n : n+sl])
		n = n + sl

		// value
		sl = binary.BigEndian.Uint32(data[n : n+4])
		n = n + 4
		if n+sl > l {
			return m, errors.New("wrong ext some keys or values are missing")
		}
		v := string(data[n : n+sl])
		n = n + sl
		m[k] = v
	}
	return m, nil
}

/*
UnPack 消息解码
消息格式
扩展数据长度|主体数据长度|扩展数据|主体数据
*/
func (m *MessagePack) UnPack(binaryMessage []byte) (Imessage, error) {
	buffer := bytes.NewReader(binaryMessage[:])

	message := &Message{
	}

	// 解析魔数 用于判断请求是否正确
	_, err := io.ReadFull(buffer, message.Header[:1])
	if err != nil {
		return nil, err
	}

	if !message.CheckMagicNumber() {
		return nil, fmt.Errorf("wrong magic number: %v", message.Header[0])
	}

	// 解析header
	_, err = io.ReadFull(buffer, message.Header[1:])
	if err != nil {
		return nil, err
	}

	// 解析消息总长度
	var totalL, serviceFuncL, extL, dataL uint32
	if err := binary.Read(buffer, binary.BigEndian, &totalL); err != nil {
		return nil, err
	}

	// 读取方法名长度
	if err := binary.Read(buffer, binary.BigEndian, &serviceFuncL); err != nil {
		return nil, err
	}

	// 读取方法名
	serviceFuncData := make([]byte, serviceFuncL)
	if l, err := io.ReadFull(buffer, serviceFuncData); l != int(serviceFuncL) || err != nil {
		return nil, fmt.Errorf("read len 0 or %w", err)
	}
	message.serviceFunc = common.SliceByteToString(serviceFuncData)

	// 读取扩展信息长度
	if err := binary.Read(buffer, binary.BigEndian, &extL); err != nil {
		return nil, err
	}

	// 读取扩展信息
	extData := make([]byte, extL)
	if l, err := io.ReadFull(buffer, extData); l != int(extL) || err != nil {
		return nil, fmt.Errorf("read len 0 or %w", err)
	}

	message.ext, err = decodeExt(extL, extData)

	// 读取信息长度
	if err := binary.Read(buffer, binary.BigEndian, &dataL); err != nil {
		return nil, err
	}

	// 读取信息
	data := make([]byte, dataL)
	if l, err := io.ReadFull(buffer, data); l != int(dataL) || err != nil {
		return nil, fmt.Errorf("read len 0 or %w", err)
	}
	message.data = data
	return message, nil
}

// func (m *MessagePack) ReadUnPack(conn net.Conn) (Imessage, error) {
// 	headData := make([]byte, common.HeaderLength)
// 	_, err := io.ReadFull(conn, headData) // ReadFull 会把msg填充满为止
// 	if err != nil {
// 		log.Println("[Read] read header error", err)
// 		return nil, err
// 	}
//
// 	header := bytes.NewReader(headData)
// 	// log.Println(string(headData))
//
// 	// 只解压head的信息，得到dataLen和msgID
// 	var extLen, dataLen uint32
// 	// 读取扩展信息长度
// 	if err := binary.Read(header, binary.BigEndian, &extLen); err != nil {
// 		return nil, err
// 	}
//
// 	// 读入消息长度
// 	if err := binary.Read(header, binary.BigEndian, &dataLen); err != nil {
// 		return nil, err
// 	}
//
// 	msg := &Message{
// 		// extLen:  extLen,
// 		// dataLen: dataLen,
// 		// count:   common.HeaderLength + extLen + dataLen + 1,
// 	}
//
// 	codecType := make([]byte, 1)
// 	{
// 		n, err := io.ReadFull(conn, codecType)
// 		if err != nil {
// 			log.Println("[Read] read codecType error", err)
// 			return nil, err
// 		}
// 		if uint32(n) != 1 {
// 			log.Println("[Read] read codecType len error")
// 			return nil, errors.New("read codecType error")
// 		}
// 	}
//
// 	extData := make([]byte, extLen)
// 	// 读取扩展信息
// 	{
// 		n, err := io.ReadFull(conn, extData)
// 		if err != nil {
// 			log.Println("[Read] read extData error", err)
// 			return nil, err
// 		}
// 		if uint32(n) != extLen {
// 			log.Println("[Read] read extData len error")
// 			return nil, errors.New("read extData error")
// 		}
// 	}
//
// 	data := make([]byte, dataLen)
// 	// 读取数据
// 	{
// 		n, err := io.ReadFull(conn, data)
// 		if err != nil {
// 			log.Println("[Read] read date error", err)
// 			return nil, err
// 		}
// 		if uint32(n) != dataLen {
// 			log.Println("[Read] read data len error")
// 			return nil, errors.New("read extData error")
// 		}
// 	}
//
// 	// 设置编码类型
// 	// msg.SetCodecType(codec.CodecType(codecType[0]))
// 	// // 获取扩展消息
// 	// msg.SetExt(extData)
// 	// 获取消息内容
// 	msg.SetData(data)
// 	return msg, nil
//
// }
