package client

import "github.com/gmsorg/gms/serialize"

type IClient interface {
	Call(serviceFunc string, request interface{}, response interface{}) error

	SetSerializeType(serializeType serialize.SerializeType) error

	Close()
}
