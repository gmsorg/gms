package connection

type IConnection interface {
	Send(reqData []byte,response interface{}) ( error)
	GetSeq() uint64
	// Read(response interface{}) error
	// Read() (protocol.Imessage, error)
}
