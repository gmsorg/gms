package protocol

/**
请求消息和返回消息体封装
*/
type Message struct {
	ExtLen  uint32
	Ext     []byte
	DataLen uint32
	Data    []byte
	Count   uint32
}

/**
初始化消息方法
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

func (m *Message) SetExtLen(extLen uint32) {
	m.ExtLen = extLen
}

func (m *Message) GetExtLen() uint32 {
	return m.ExtLen
}

func (m *Message) SetExt(ext []byte) {
	m.Ext = ext
}

func (m *Message) GetExt() []byte {
	return m.Ext
}

func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) GetData() []byte {
	return m.Data
}
