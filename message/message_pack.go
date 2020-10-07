package message

/**
消息编码、解码
*/
type MessagePack struct {
}

/**
消息编码
*/
func (m *MessagePack) Encode(msg Imessage) ([]byte, error) {
	panic("implement me")
}

/**
消息解码
*/
func (m *MessagePack) Decode(bytes []byte) (Imessage, error) {
	panic("implement me")
}
