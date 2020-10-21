package codec

type ICodec interface {
	Encode(i interface{}) ([]byte, error)
	Decode(data []byte, i interface{}) error
}

type CodecType byte

const (
	JSON CodecType = iota
	Msgpack
	Gob
)

var (
	Codecs = map[CodecType]ICodec{
		JSON:    &JsonCodec{},
		Msgpack: &MsgpackCode{},
		Gob:     &GobCodec{},
	}
)

func GetCodec(codecType CodecType) ICodec {
	if codec, ok := Codecs[codecType]; ok {
		return codec
	}
	return nil
}
