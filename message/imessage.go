package message

/**
消息编码、解码
*/
type IMessage interface {
	Encode(msg IMessage) ([]byte, error) // 消息编码方法
	Decode([]byte) (IMessage, error)     // 消息解码方法
}
