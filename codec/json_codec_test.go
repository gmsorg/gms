package codec

import (
	"fmt"
	"testing"
)

func TestJsonCodec_Encode(t *testing.T) {
	jsonCodec := JsonCodec{}
	demoEncode := demo{
		Str:    "CrazyWolf",
		Number: 20,
		Inter:  1603024853073143124,
	}
	ub, err := jsonCodec.Encode(demoEncode)
	if err != nil {
		fmt.Println(err)
	}

	demoDecode := &demo{}
	err = jsonCodec.Decode(ub, demoDecode)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(demoDecode.Inter)
	fmt.Println(demoEncode.Inter)
	fmt.Println(demoEncode.Str == demoDecode.Str, demoEncode.Number == demoDecode.Number, demoEncode == *demoDecode)
}
