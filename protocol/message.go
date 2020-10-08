package protocol

/*
Message 请求消息和返回消息体封装
*/
type Message struct {
	ExtLen  uint32
	Ext     []byte
	DataLen uint32
	Data    []byte
	Count   uint32
}

/*
NewMessage 初始化消息方法
*/
func NewMessage(ext, data []byte) Imessage {
	return &Message{
		ExtLen:  uint32(len(ext)),
		Ext:     ext,
		DataLen: uint32(len(data)),
		Data:    data,
		Count:   uint32(len(ext)) + uint32(len(data)),
	}
}

/*
SetExtLen 设置扩展信息长度
*/
func (m *Message) SetExtLen(extLen uint32) {
	m.ExtLen = extLen
}

/*
GetExtLen 获取扩展数据的长度
*/
func (m *Message) GetExtLen() uint32 {
	return m.ExtLen
}

/*
SetExt 设置扩展数据
*/
func (m *Message) SetExt(ext []byte) {
	m.Ext = ext
}

/*
GetExt 获取扩展数据
*/
func (m *Message) GetExt() []byte {
	return m.Ext
}

/*
SetDataLen 设置主体数据段长度
*/
func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

/*
GetDataLen 获取主体数据段长度
*/
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

/*
SetData 设置主体数据内容
*/
func (m *Message) SetData(data []byte) {
	m.Data = data
}

/*
GetData 获取主体数据内容
*/
func (m *Message) GetData() []byte {
	return m.Data
}
