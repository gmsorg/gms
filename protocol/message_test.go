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
	fmt.Println("version:", message.GetVersion())

	message.SetMessageType(Heartbeat)
	fmt.Println("GetMessageType Heartbeat", message.GetMessageType())
	message.SetMessageType(Request)
	fmt.Println("GetMessageType Request", message.GetMessageType())
	message.SetMessageType(Response)
	fmt.Println("GetMessageType Response", message.GetMessageType())
	message.SetMessageType(ResponseError)
	fmt.Println("GetMessageType ResponseError", message.GetMessageType())

	seq := rand.Int63n(1000000)
	fmt.Println("set seq:", seq)
	message.SetSeq(uint64(seq))
	fmt.Println("get seq:", message.GetSeq())

	message.SetExt(nil)
	message.SetData([]byte("hello"))



}
