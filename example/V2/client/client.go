package main

import (
	"fmt"
	"time"

	"github.com/akkagao/gms/client"
	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/example/V1/vo"
)

/*
	模拟客户端
*/
func main() {
	start := time.Now()
	discovery := discovery.NewP2PDiscovery("127.0.0.1:9000")

	demoClient, err := client.NewClient(discovery)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := vo.AddUserReq{Name: "aaa"}
	res := &vo.AddUserRes{}

	err = demoClient.Call("user.Add", req, res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	fmt.Println(time.Since(start))
	// conn.Read()

}
