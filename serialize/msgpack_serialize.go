package serialize

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)


type MsgpackSerialize struct{}


func (c *MsgpackSerialize) Serialize(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := enc.Encode(i)
	return buf.Bytes(), err
}

func (c *MsgpackSerialize) UnSerialize(data []byte, i interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(i)
	return err
}
