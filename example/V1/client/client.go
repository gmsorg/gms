package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/example/V1/vo"
	"github.com/akkagao/gms/protocol"
)

/*
	模拟客户端
*/
func main() {
	start := time.Now()

	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求，给服务端开启服务的机会
	// time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	// bb := common.NewBytePoolCap(20, 120, 120)

	wg := sync.WaitGroup{}

	cout := 10
	wg.Add(cout)
	for i := 0; i < cout; i++ {
		// go func(i int) {
		// fmt.Println(i)
		_, err = conn.Write(funcName(i))
		if err != nil {
			fmt.Println(err)
		}

		// buf := [200]byte{}

		buf := common.BytePool.Get()

		n, err := conn.Read(buf[0:])
		fmt.Println(string(buf[:n]))

		common.BytePool.Put(buf)

		if err != nil {
			if err == io.EOF {
				wg.Done()
				return
			}
			return
			wg.Done()
		}

		// time.Sleep(time.Second)
		wg.Done()
		// }(i)

	}
	wg.Wait()
	// select {}

	fmt.Println(time.Since(start))
	// conn.Read()

}

func funcName(i int) []byte {
	addUser := vo.AddUserReq{
		Name: fmt.Sprintf("%v--%v--%v", i, "hello", common.GenIdentity()),
	}

	addUserData, err := json.Marshal(addUser)
	if err != nil {
		fmt.Println(err)
	}

	message := protocol.NewMessage([]byte("user.Add"), addUserData)
	mp := protocol.MessagePack{}
	encodeMessage, err := mp.Pack(message)
	if err != nil {
		fmt.Println(err)
	}
	return encodeMessage
}
