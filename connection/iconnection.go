package connection

import "github.com/gmsorg/gms/protocol"

type IConnection interface {
	Send(reqData []byte) (protocol.Imessage, error)
	GetSeq() uint64
	// Read(response interface{}) error
	// Read() (protocol.Imessage, error)
}
