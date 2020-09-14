package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/akka/gms/common"
	"github.com/akka/gms/example/user"
)

type gmsConnection struct {
	conn net.Conn
}

func (c *gmsConnection) Call() {
	// Done := c.conn.Go(ctx, servicePath, serviceMethod, args, reply, make(chan *Call, 1)).Done
	// {"service_name":"UserServiceImpl","method_name":"GetUser"}

	r := user.GetUserReq{
		Id: time.Now().UnixNano() / 10e6,
	}
	rb, _ := json.Marshal(r)

	reqMessage := common.ReqMessage{
		ServiceName: "UserServiceImpl",
		MethodName:  "GetUser",
		ReqData:     rb,
	}

	b, err := json.Marshal(reqMessage)
	if err != nil {
		log.Fatalln(err)
	}

	c.conn.Write(b)

	// --------------read-----------
	result := bytes.NewBuffer(nil)

	var buf [512]byte
	for {
		n, err := c.conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		if n < 512 {
			break
		}
	}

	res := &user.GetUserRes{}
	resMessage := common.ResMessage{}
	// resultValues := []reflect.Value{}
	fmt.Println("====result.Bytes======")
	fmt.Println(string(result.Bytes()))
	fmt.Println("======result.Bytes====")
	json.Unmarshal(result.Bytes(), &resMessage)

	fmt.Println(resMessage.Code, resMessage.Msg)

	fmt.Println("==========resMessage.ResData=========")
	fmt.Println(resMessage.ResData)
	fmt.Println("===========resMessage.ResData========")
	json.Unmarshal(resMessage.ResData, res)

	fmt.Println(res)

}

func Dial(address string) (*gmsConnection, error) {
	gmsConn := &gmsConnection{}
	conn, err := net.DialTimeout("tcp", address, time.Second*3)
	if err != nil {
		log.Fatal(err)
	}

	gmsConn.conn = conn
	return gmsConn, nil
}
