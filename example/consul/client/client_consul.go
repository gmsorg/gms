package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gmsorg/gms/client"
	"github.com/gmsorg/gms/codec"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/example/model"
)

/*
	模拟客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery, err := discovery.NewConsulDiscover("gmsDemo", []string{"127.0.0.1:8500"})
	if err != nil {
		log.Fatal(err)
	}

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	req := &model.AdditionReq{NumberA: 10, NumberB: 20}

	for {
		time.Sleep(time.Second)
		// 接收返回值的对象
		res := &model.AdditionRes{}
		// 调用服务
		err = additionClient.Call("addition", req, res)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(fmt.Sprintf("%d+%d=%d", req.NumberA, req.NumberB, res.Result))
	}

}
