package protocol

import (
	"github.com/akkagao/gms/codec"
	"github.com/akkagao/gms/common"
)

/*
Message 请求消息和返回消息体封装
*/
type Message struct {
	extLen    uint32
	ext       []byte
	codecType codec.CodecType
	dataLen   uint32
	data      []byte
	count     uint32
}

/*
NewMessage 初始化消息方法
*/
func NewMessage(ext, data []byte, codecType codec.CodecType) Imessage {
	return &Message{
		extLen:    uint32(len(ext)),
		ext:       ext,
		codecType: codecType,
		dataLen:   uint32(len(data)),
		data:      data,
		// count= 消息头长度（描述扩展信息和消息数据的长度信息）+ 扩展信息的长度+消息长度+编码方式信息长度（1一个字节）
		count: common.HeaderLength + uint32(len(ext)) + uint32(len(data)) + 1,
	}
}

/*
SetExtLen 设置扩展信息长度
*/
func (m *Message) SetExtLen(extLen uint32) {
	m.extLen = extLen
}

/*
GetExtLen 获取扩展数据的长度
*/
func (m *Message) GetExtLen() uint32 {
	return m.extLen
}

/*
SetExt 设置扩展数据
*/
func (m *Message) SetExt(ext []byte) {
	m.ext = ext
}

/*
GetExt 获取扩展数据
*/
func (m *Message) GetExt() []byte {
	return m.ext
}

// 设置编码方式
func (m *Message) SetCodecType(codecType codec.CodecType) {
	m.codecType = codecType
}

// 获取编码方式
func (m *Message) GetCodecType() codec.CodecType {
	return m.codecType
}

/*
SetDataLen 设置主体数据段长度
*/
func (m *Message) SetDataLen(dataLen uint32) {
	m.dataLen = dataLen
}

/*
GetDataLen 获取主体数据段长度
*/
func (m *Message) GetDataLen() uint32 {
	return m.dataLen
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

/*
GetCount 获取消息总长度
*/
func (m *Message) GetCount() uint32 {
	return m.count
}
