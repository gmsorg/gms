package client

type IClient interface {
	Call(serviceFunc string, request interface{}) (response interface{}, err error)
	Close()
}
