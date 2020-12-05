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

func GetSerialize(serializeType SerializeType) ISerialize {
	if serialize, ok := Serializes[serializeType]; ok {
		return serialize
	}
	return nil
}
