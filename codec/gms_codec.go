package codec

type GmsCodec struct {
}

func NewGmsCodec() ICodec {
	return &GmsCodec{}
}
func (g *GmsCodec) EnCodec(i interface{}) ([]byte, error) {
	panic("implement me")
}

func (g *GmsCodec) DeCodec(data []byte) chan []byte {
	panic("implement me")
}
