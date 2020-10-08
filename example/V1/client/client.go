package main

import (
	"encoding/json"
	"fmt"
	"net"
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

	_, err = conn.Write(encodeMessage)
	if err != nil {
		fmt.Println(err)
	}

	// conn.Read()

}
