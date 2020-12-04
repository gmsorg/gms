package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/akkagao/gms/client"
	"github.com/akkagao/gms/codec"
	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/example/model"
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
	// todo Msgpack 发送到服务端会解析错误。header 解析失败了(明天继续)
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	start := time.Now()

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		waitGroup.Add(1)
		go func(i int) {
			for j := 0; j < 1; j++ {
				rand.Seed(time.Now().UnixNano())
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
