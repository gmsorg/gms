package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/akka/gms/common"
)

/*
MessagePack 消息编码、解码
实现gnet.ICodec 接口
*/
type MessagePack struct {
	// Message
}

/*
Encode 消息编码
消息格式
扩展数据长度|主体数据长度|扩展数据|主体数据
*/
func (m *MessagePack) Encode(message Imessage) ([]byte, error) {

	fmt.Println(message.GetExtLen(), message.GetDataLen())

	result := make([]byte, 0)

	buffer := bytes.NewBuffer(result)

	if err := binary.Write(buffer, binary.BigEndian, message.GetExtLen()); err != nil {
		s := fmt.Sprintf("[Encode] Pack ExtLen error , %v", err)
		return nil, errors.New(s)
	}

	if err := binary.Write(buffer, binary.BigEndian, message.GetDataLen()); err != nil {
		s := fmt.Sprintf("[Encode] Pack DataLen error , %v", err)
		return nil, errors.New(s)
	}

	if message.GetExtLen() > 0 {
		if err := binary.Write(buffer, binary.BigEndian, message.GetExt()); err != nil {
			s := fmt.Sprintf("[Encode] Pack Ext error , %v", err)
			return nil, errors.New(s)
		}
	}

	if message.GetDataLen() > 0 {
		if err := binary.Write(buffer, binary.BigEndian, message.GetData()); err != nil {
			s := fmt.Sprintf("[Encode] Pack Data error , %v", err)
			return nil, errors.New(s)
		}
	}

	fmt.Println(string(buffer.Bytes()))
	return buffer.Bytes(), nil
}

/*
Decode 消息解码
消息格式
扩展数据长度|主体数据长度|扩展数据|主体数据
*/
func (m *MessagePack) Decode(binaryMessage []byte) (Imessage, error) {
	header := bytes.NewReader(binaryMessage[:common.HeaderLength])

	// 只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	// 读取扩展信息长度
	if err := binary.Read(header, binary.BigEndian, &msg.ExtLen); err != nil {
		return nil, err
	}

	// 读入消息长度
	if err := binary.Read(header, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 截取消息投后的所有内容
	content := binaryMessage[common.HeaderLength:]
	// 获取扩展消息
	msg.Ext = content[:msg.ExtLen]
	// 获取消息内容
	msg.Data = content[msg.ExtLen:]

	// extBuff := bytes.NewReader(content[:msg.ExtLen])
	// contentBuff := bytes.NewReader(content[msg.ExtLen:])

	// // 读取扩展信息
	// if err := binary.Read(extBuff, binary.BigEndian, msg.Ext); err != nil {
	// 	return nil, err
	// }
	//
	// if err := binary.Read(contentBuff, binary.BigEndian, msg.Data); err != nil {
	// 	return nil, err
	// }
	// 这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil

}
