package message

/**
请求消息和返回消息体封装
*/
type Message struct {
	extLen  int
	ext     []byte
	dataLen int
	data    []byte
}

/**
初始化消息方法
*/
func NewMessage(ext, data []byte) Imessage {
	return &Message{
		extLen:  len(ext),
		ext:     ext,
		dataLen: len(data),
		data:    data,
	}
}

func (m *Message) SetExtLen(extLen int) {
	m.extLen = extLen
}

func (m *Message) GetExtLen() int {
	return m.extLen
}

func (m *Message) SetExt(ext []byte) {
	m.ext = ext
}

func (m *Message) GetExt() []byte {
	return m.ext
}

func (m *Message) SetDataLen(dataLen int) {
	m.dataLen = dataLen
}

func (m *Message) GetDataLen() int {
	return m.dataLen
}

func (m *Message) SetData(data []byte) {
	m.data = data
}

func (m *Message) GetData() []byte {
	return m.data
}
