package codec

import (
	"fmt"
	"testing"
)

type demo struct {
	Str    string
	Number int
	Inter  interface{}
}

func TestMsgpackCodec_Encode(t *testing.T) {
	msgpackCodec := MsgpackCode{}
	demoEncode := demo{
		Str:    "CrazyWolf",
		Number: 20,
		Inter:  1603024853073143124,
	}
	ub, err := msgpackCodec.Encode(demoEncode)
	if err != nil {
		fmt.Println(err)
	}

	demoDecode := &demo{}
	err = msgpackCodec.Decode(ub, demoDecode)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(demoDecode.Inter)
	fmt.Println(demoEncode.Inter)
	fmt.Println(demoEncode.Str == demoDecode.Str, demoEncode.Number == demoDecode.Number, demoEncode == *demoDecode)
}
