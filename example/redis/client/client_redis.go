package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gmsorg/gms/client"
	"github.com/gmsorg/gms/serialize"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/example/model"
)

/*
	模拟客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery, err := discovery.NewRedisDiscover("gmsDemo", "127.0.0.1:6379")
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
	additionClient.SetSerializeType(serialize.Msgpack)

	// 请求对象
	// req := &model.AdditionReq{NumberA: 10, NumberB: 20}

	start := time.Now()

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			for j := 0; j < 1000; j++ {
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
