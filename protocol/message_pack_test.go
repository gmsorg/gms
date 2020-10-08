package protocol

import (
	"fmt"
	"testing"
)

func TestMessagePack_Encode(t *testing.T) {
	message := NewMessage([]byte("user.add"), []byte("hello"))
	mp := MessagePack{}
	encodeData, err := mp.Encode(message)
	fmt.Println(string(encodeData), err)
}

func TestMessagePack_Decode(t *testing.T) {
	message := NewMessage([]byte("user.add"), []byte("hello world"))
	mp := MessagePack{}
	encodeData, err := mp.Encode(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encodeData))
	m, err := mp.Decode(encodeData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GetExtLen:", m.GetExtLen(), "GetDataLen:", m.GetDataLen(),
		"GetExt:", string(m.GetExt()), "GetData:", string(m.GetData()))
}
