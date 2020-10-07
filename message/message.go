package message

/**
消息编码、解码
*/
type Message struct {
}

func (m *Message) Encode(msg IMessage) ([]byte, error) {
	panic("implement me")
}

func (m *Message) Decode(bytes []byte) (IMessage, error) {
	panic("implement me")
}
