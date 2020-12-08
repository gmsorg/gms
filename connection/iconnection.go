package connection

import "github.com/gmsorg/gms/protocol"

type IConnection interface {
	Send(reqData []byte, response interface{}) error
	SendM(message protocol.Imessage, response interface{}) error
	SetConnId(connId int)
	GetSeq() uint64
	// Read(response interface{}) error
	// Read() (protocol.Imessage, error)
}
