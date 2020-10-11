package connection

type IConnection interface {
	Send([]byte) error
}
