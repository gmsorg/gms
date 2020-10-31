package main

import (
	"fmt"

	"github.com/akkagao/gms/client"
	"github.com/akkagao/gms/codec"
	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/example/model"
)

/*
	模拟客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery := discovery.NewP2PDiscovery("127.0.0.1:1024")

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	req := &model.AdditionReq{NumberA: 10, NumberB: 20}
	// 接收返回值的对象
	res := &model.AdditionRes{}

	// 调用服务
	err = additionClient.Call("addition", req, res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%d+%d=%d", req.NumberA, req.NumberB, res.Result))
}
