package protocol

import (
	"log"
	"testing"

	"github.com/akkagao/gms/codec"
)

func TestMessagePack_Encode(t *testing.T) {
	message := NewMessage([]byte("user.add"), []byte("hello"), codec.JSON)
	mp := MessagePack{}
	encodeData, err := mp.Pack(message)
	log.Println(string(encodeData), err)
}

func TestMessagePack_Decode(t *testing.T) {
	message := NewMessage([]byte("user.add"), []byte("hello world"), codec.Msgpack)
	mp := MessagePack{}
	encodeData, err := mp.Pack(message)
	if err != nil {
		log.Println(err)
	}
	log.Println("encodeData:", string(encodeData))
	m, err := mp.UnPack(encodeData)
	if err != nil {
		log.Println(err)
	}
	log.Println("GetExtLen:", m.GetExtLen(), "GetDataLen:", m.GetDataLen(),
		"GetCodecType:", m.GetCodecType(), " GetExt:", string(m.GetExt()), "GetData:", string(m.GetData()))
	log.Println("GetCount:", m.GetCount())
}
