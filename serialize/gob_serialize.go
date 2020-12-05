package serialize

import (
	"bytes"
	"encoding/gob"
)

type GobSerialize struct {
}

func (c *GobSerialize) UnSerialize(data []byte, i interface{}) error {
	enc := gob.NewDecoder(bytes.NewBuffer(data))
	err := enc.Decode(i)
	return err
}

func (c *GobSerialize) Serialize(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	return buf.Bytes(), err
}
