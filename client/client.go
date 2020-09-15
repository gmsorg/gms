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

var size = 4
var width = 10
var capWidth = 16

type GmsConnection struct {
	conn    net.Conn
	bufPool *common.BytePoolCap
}

/**
创建连接
*/
func Dial(address string) (*GmsConnection, error) {
	var size = 50
	var width = 512
	var capWidth = 512

	gmsConn := &GmsConnection{
		bufPool: common.NewBytePoolCap(size, width, capWidth),
	}

	conn, err := net.DialTimeout("tcp", address, time.Second*3)
	if err != nil {
		log.Fatal(err)
	}

	gmsConn.conn = conn
	return gmsConn, nil
}

func (gc *GmsConnection) CommonCall(serviceName, methodName string, req, res interface{}) {
	rb, _ := json.Marshal(req)

	reqMessage := common.ReqMessage{
		ServiceName: serviceName,
		MethodName:  methodName,
		ReqData:     rb,
	}

	b, err := json.Marshal(reqMessage)
	if err != nil {
		log.Fatalln(err)
	}

	gc.conn.Write(b)

	result := bytes.NewBuffer(nil)

	// var buf [512]byte

	buf := gc.bufPool.Get()
	defer gc.bufPool.Put(buf)
	for {
		n, err := gc.conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		if n < gc.bufPool.WidthCap() {
			break
		}
	}

	// res := &user.GetUserRes{}
	resMessage := common.ResMessage{}
	// resultValues := []reflect.Value{}
	// fmt.Println("====result.Bytes======")
	// fmt.Println(string(result.Bytes()))
	// fmt.Println("======result.Bytes====")
	json.Unmarshal(result.Bytes(), &resMessage)

	// fmt.Println(resMessage.Code, resMessage.Msg)
	//
	// fmt.Println("==========resMessage.ResData=========")
	// fmt.Println(resMessage.ResData)
	// fmt.Println("===========resMessage.ResData========")
	json.Unmarshal(resMessage.ResData, res)

	// fmt.Println(res)

}

func (gc *GmsConnection) Call() {
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

	gc.conn.Write(b)

	// --------------read-----------
	result := bytes.NewBuffer(nil)

	// var buf [512]byte

	buf := gc.bufPool.Get()
	defer gc.bufPool.Put(buf)
	for {
		n, err := gc.conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		if n < gc.bufPool.WidthCap() {
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
