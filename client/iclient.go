package client

import "github.com/gmsorg/gms/codec"

type IClient interface {
	Call(serviceFunc string, request interface{}, response interface{}) error

	SetCodecType(codecType codec.CodecType) error

	Close()
}
