package codec

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

// JsonCodec uses messagepack marshaler and unmarshaler.
type MsgpackCode struct{}

// Encode encodes an object into slice of bytes.
func (c *MsgpackCode) Encode(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	// enc.UseJSONTag(true)
	err := enc.Encode(i)
	return buf.Bytes(), err
}

// Decode decodes an object from slice of bytes.
func (c *MsgpackCode) Decode(data []byte, i interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	// dec.UseJSONTag(true)
	err := dec.Decode(i)
	return err
}
