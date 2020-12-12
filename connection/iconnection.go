package connection

import "github.com/gmsorg/gms/protocol"

type IConnection interface {
	Send(message protocol.Imessage, response interface{}) error

	SetConnId(connId int)
}
