package serialize

import (
	"log"
	"testing"
)

type demo struct {
	Str    string
	Number int
	Inter  interface{}
}

func TestMsgpackCodec_Encode(t *testing.T) {
	msgpackCodec := MsgpackSerialize{}
	demoEncode := demo{
		Str:    "CrazyWolf",
		Number: 20,
		Inter:  1603024853073143124,
	}
	ub, err := msgpackCodec.Serialize(demoEncode)
	if err != nil {
		log.Println(err)
	}

	demoDecode := &demo{}
	err = msgpackCodec.UnSerialize(ub, demoDecode)
	if err != nil {
		log.Println(err)
	}

	log.Println(demoDecode.Inter)
	log.Println(demoEncode.Inter)
	log.Println(demoEncode.Str == demoDecode.Str, demoEncode.Number == demoDecode.Number, demoEncode == *demoDecode)
}
