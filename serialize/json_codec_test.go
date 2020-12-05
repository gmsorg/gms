package serialize

import (
	"log"
	"testing"
)

func TestJsonCodec_Encode(t *testing.T) {
	jsonCodec := JsonSerialize{}
	demoEncode := demo{
		Str:    "CrazyWolf",
		Number: 20,
		Inter:  1603024853073143124,
	}
	ub, err := jsonCodec.Serialize(demoEncode)
	if err != nil {
		log.Println(err)
	}

	demoDecode := &demo{}
	err = jsonCodec.UnSerialize(ub, demoDecode)
	if err != nil {
		log.Println(err)
	}

	log.Println(demoDecode.Inter)
	log.Println(demoEncode.Inter)
	log.Println(demoEncode.Str == demoDecode.Str, demoEncode.Number == demoDecode.Number, demoEncode == *demoDecode)
}
