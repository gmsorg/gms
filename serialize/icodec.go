package serialize

type ISerialize interface {
	Serialize(i interface{}) ([]byte, error)
	UnSerialize(data []byte, i interface{}) error
}

type SerializeType byte

const (
	JSON SerializeType = iota
	Msgpack
	Gob
)

var (
	Serializes = map[SerializeType]ISerialize{
		JSON:    &JsonSerialize{},
		Msgpack: &MsgpackSerialize{},
		Gob:     &GobSerialize{},
	}
)

func GetCodec(codecType SerializeType) ISerialize {
	if codec, ok := Serializes[codecType]; ok {
		return codec
	}
	return nil
}
