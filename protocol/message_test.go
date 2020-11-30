package protocol

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	message := NewMessage()
	fmt.Println("CheckMagicNumber:", message.CheckMagicNumber())

	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(127)
	fmt.Println("set version:", v)
	message.SetVersion(byte(v))
	fmt.Println("version:", message.Version())

	message.SetMessageType(Heartbeat)
	fmt.Println("MessageType Heartbeat", message.MessageType())
	message.SetMessageType(Request)
	fmt.Println("MessageType Request", message.MessageType())
	message.SetMessageType(Response)
	fmt.Println("MessageType Response", message.MessageType())
	message.SetMessageType(ResponseError)
	fmt.Println("MessageType ResponseError", message.MessageType())

	seq := rand.Int63n(1000000)
	fmt.Println("set seq:", seq)
	message.SetSeq(uint64(seq))
	fmt.Println("get seq:", message.Seq())

	message.SetExt(nil)
	message.SetData([]byte("hello"))



}
