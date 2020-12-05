package serialize

import (
	"bytes"
	"encoding/json"
)

type JsonSerialize struct {
}

func (j *JsonSerialize) Serialize(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func (j *JsonSerialize) UnSerialize(data []byte, i interface{}) error {
	d := json.NewDecoder(bytes.NewBuffer(data))
	// 解决json返序列化，interface{}接收。数字被解析为float64,精度丢失问题
	d.UseNumber()
	return d.Decode(&i)
	// return json.Unmarshal(data, i)
}
