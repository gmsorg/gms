package connection

type IConnection interface {
	Send(reqData []byte) error
	Read(response interface{}) error
}
