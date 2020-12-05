package serialize

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

// JsonSerialize uses messagepack marshaler and unmarshaler.
type MsgpackSerialize struct{}

// Serialize encodes an object into slice of bytes.
func (c *MsgpackSerialize) Serialize(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	// enc.UseJSONTag(true)
	err := enc.Encode(i)
	return buf.Bytes(), err
}

// UnSerialize decodes an object from slice of bytes.
func (c *MsgpackSerialize) UnSerialize(data []byte, i interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	// dec.UseJSONTag(true)
	err := dec.Decode(i)
	return err
}
