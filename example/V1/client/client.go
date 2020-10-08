package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/akka/gms/example/V1/vo"
	"github.com/akka/gms/protocol"
)

/*
	模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	addUser := vo.AddUserReq{
		Name: "hello",
	}

	addUserData, err := json.Marshal(addUser)
	if err != nil {
		fmt.Println(err)
	}

	message := protocol.NewMessage([]byte("user.Add"), addUserData)
	mp := protocol.MessagePack{}
	encodeMessage, err := mp.Encode(message)
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		_, err = conn.Write(encodeMessage)
		if err != nil {
			fmt.Println(err)
		}

		buf := [512]byte{}

		n, err := conn.Read(buf[0:])
		fmt.Println(string(buf[:n]))
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}

		// time.Sleep(time.Second)
		wg.Done()
	}
	wg.Wait()
	select {}

	// conn.Read()

}
