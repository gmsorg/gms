package protocol

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMessagePack_Encode(t *testing.T) {
	message := newMessage()
	mp := MessagePack{}
	encodeData, err := mp.Pack(message)
	fmt.Println(encodeData)
	fmt.Println(string(encodeData), err)
}

func TestMessagePack_Decode(t *testing.T) {
	message := newMessage()
	mp := MessagePack{}
	encodeData, err := mp.Pack(message)
	fmt.Println(err)
	fmt.Println("===============")
	m, err := mp.UnPack(encodeData)
	fmt.Println(err)
	// fmt.Println(m.GetHeader())

	fmt.Println("version:", m.GetVersion())
	fmt.Println("GetMessageType:", m.GetMessageType())
	fmt.Println("GetCompressType ", m.GetCompressType())
	fmt.Println("get seq:", m.GetSeq())
	fmt.Println("GetServiceFunc:", m.GetServiceFunc())
	fmt.Println("ext:", m.GetExt())
	fmt.Println("data:", string(m.GetData()))

}

func newMessage() Imessage {
	message := NewMessage()
	// fmt.Println("CheckMagicNumber:", message.CheckMagicNumber())

	rand.Seed(time.Now().UnixNano())

	v := rand.Intn(127)
	// fmt.Println("set version:", v)
	message.SetVersion(byte(v))
	// fmt.Println("version:", message.GetVersion())

	// message.SetMessageType(Heartbeat)
	// fmt.Println("GetMessageType Heartbeat", message.GetMessageType())
	message.SetMessageType(Request)
	// fmt.Println("GetMessageType Request", message.GetMessageType())
	// message.SetMessageType(Response)
	// fmt.Println("GetMessageType Response", message.GetMessageType())
	// message.SetMessageType(ResponseError)
	// fmt.Println("GetMessageType ResponseError", message.GetMessageType())

	// message.SetCompressType(Gzip)
	message.SetCompressType(None)
	// fmt.Println("GetCompressType ", message.GetCompressType())

	seq := rand.Int63n(1000000)
	// fmt.Println("set seq:", seq)
	message.SetSeq(uint64(seq))
	// fmt.Println("get seq:", message.GetSeq())

	message.SetExt(map[string]string{"id": "aa"})
	message.SetData([]byte("hello"))
	message.SetServiceFunc("GetServiceFunc")
	return message
}
