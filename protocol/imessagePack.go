package protocol

/*
IMessagePack 消息封包、解包
*/
type IMessagePack interface {
	// 请求消息封包方法
	Pack(msg Imessage) ([]byte, error)
	// 请求消息解包方法
	UnPack([]byte) (Imessage, error)
	//	从conn中读取数据解包
	// ReadUnPack(net.Conn) (Imessage, error)
}
