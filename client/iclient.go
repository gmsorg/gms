package client

type IClient interface {
	Call(serviceFunc string, request interface{}, response interface{}) error
	Close()
}
