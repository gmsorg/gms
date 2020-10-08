package protocol

/**
消息编码、解码
*/
type IMessagePack interface {
	Encode(msg Imessage) ([]byte, error) // 请求消息编码方法
	Decode([]byte) (Imessage, error)     // 请求消息解码方法
}
