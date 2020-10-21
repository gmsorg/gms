package connection

import "github.com/akkagao/gms/protocol"

type IConnection interface {
	Send(reqData []byte) error
	// Read(response interface{}) error
	Read() (protocol.Imessage, error)
}
