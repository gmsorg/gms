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

	fmt.Println("version:", m.Version())
	fmt.Println("MessageType:", m.MessageType())
	fmt.Println("CompressType ", m.CompressType())
	fmt.Println("get seq:", m.Seq())
	fmt.Println("ServiceFunc:", m.ServiceFunc())
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
	// fmt.Println("version:", message.Version())

	// message.SetMessageType(Heartbeat)
	// fmt.Println("MessageType Heartbeat", message.MessageType())
	message.SetMessageType(Request)
	// fmt.Println("MessageType Request", message.MessageType())
	// message.SetMessageType(Response)
	// fmt.Println("MessageType Response", message.MessageType())
	// message.SetMessageType(ResponseError)
	// fmt.Println("MessageType ResponseError", message.MessageType())

	// message.SetCompressType(Gzip)
	message.SetCompressType(None)
	// fmt.Println("CompressType ", message.CompressType())

	seq := rand.Int63n(1000000)
	// fmt.Println("set seq:", seq)
	message.SetSeq(uint64(seq))
	// fmt.Println("get seq:", message.Seq())

	message.SetExt(map[string]string{"id": "aa"})
	message.SetData([]byte("hello"))
	message.SetServiceFunc("ServiceFunc")
	return message
}
