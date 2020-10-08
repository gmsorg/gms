package protocol

/*
IMessagePack 消息编码、解码
*/
type IMessagePack interface {
	// 请求消息编码方法
	Encode(msg Imessage) ([]byte, error)
	// 请求消息解码方法
	Decode([]byte) (Imessage, error)
}
