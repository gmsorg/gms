package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gmsorg/gms/client"
	"github.com/gmsorg/gms/codec"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/example/model"
)

/*
客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery := discovery.NewP2PDiscover([]string{"127.0.0.1:1024"})

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		log.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	start := time.Now()

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			for j := 0; j < 10; j++ {
				rand.Seed(time.Now().UnixNano())
				// req := &model.AdditionReq{NumberA: 100, NumberB: 200}
				req := &model.AdditionReq{NumberA: rand.Intn(100), NumberB: rand.Intn(200)}

				// 接收返回值的对象
				res := &model.AdditionRes{}

				// 调用服务
				err = additionClient.Call("addition", req, res)
				if err != nil {
					log.Println(err)
				}
				log.Println(fmt.Sprintf("%v-%v :%d+%d=%d", i, j, req.NumberA, req.NumberB, res.Result))
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	fmt.Println(time.Since(start))
}
