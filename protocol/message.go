package protocol

import (
	"encoding/binary"

	"github.com/akkagao/gms/codec"
)

/*
Message 请求消息和返回消息体封装
*/
type Message struct {
	Header
	serviceFunc string
	ext         map[string]string
	data        []byte
}

/*
NewMessage 初始化消息方法
*/
func NewMessage() Imessage {
	header := Header{}

	message := &Message{
		Header: header,
		// ext:    make(map[string]string),
	}
	message.setMagicNumber()

	return message
}

func (m *Message) GetHeader() Header {
	return m.Header
}

func (m *Message) setMagicNumber() {
	m.Header[0] = magicNumber
}

// CheckMagicNumber checks whether header starts rpcx magic number.
func (m *Message) CheckMagicNumber() bool {
	return m.Header[0] == magicNumber
}

// GetVersion returns version of rpcx protocol.
func (m *Message) GetVersion() byte {
	return m.Header[1]
}

// SetVersion sets version for this header.
func (m *Message) SetVersion(v byte) {
	m.Header[1] = v
}

// GetMessageType returns the message type.
func (m *Message) GetMessageType() MessageType {
	return MessageType(m.Header[2])
}

// SetMessageType sets message type.
func (m *Message) SetMessageType(mt MessageType) {
	m.Header[2] = byte(mt)
}

// GetCompressType returns compression type of messages.
func (m *Message) GetCompressType() CompressType {
	return CompressType((m.Header[3]))
}

// SetCompressType sets the compression type.
func (m *Message) SetCompressType(ct CompressType) {
	m.Header[3] = byte(ct)
}

// GetSerializeType returns serialization type of payload.
func (m *Message) GetSerializeType() codec.CodecType {
	return codec.CodecType((m.Header[4]))
}

// SetSerializeType sets the serialization type.
func (m *Message) SetSerializeType(ct codec.CodecType) {
	m.Header[4] = byte(ct)
}

// GetSeq returns sequence number of messages.
func (m *Message) GetSeq() uint64 {
	return binary.BigEndian.Uint64(m.Header[5:])
}

// SetSeq sets  sequence number.
func (m *Message) SetSeq(seq uint64) {
	binary.BigEndian.PutUint64(m.Header[5:], seq)
}

// 获取请求方法名
func (m *Message) GetServiceFunc() string {
	return m.serviceFunc
}

// 设置请求方法名
func (m *Message) SetServiceFunc(serviceFunc string) {
	m.serviceFunc = serviceFunc
}

/*
SetExt 设置扩展数据
*/
func (m *Message) SetExt(ext map[string]string) {
	m.ext = ext
}

/*
GetExt 获取扩展数据
*/
func (m *Message) GetExt() map[string]string {
	return m.ext
}

/*
SetData 设置主体数据内容
*/
func (m *Message) SetData(data []byte) {
	m.data = data
}

/*
GetData 获取主体数据内容
*/
func (m *Message) GetData() []byte {

	return m.data
}
