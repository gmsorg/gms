package codec

type ICodec interface {
	EnCodec(i interface{}) ([]byte, error)
	DeCodec(data []byte) chan []byte
}
