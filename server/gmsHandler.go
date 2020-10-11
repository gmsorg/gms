package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/satori/go.uuid"

	"github.com/akkagao/gms/protocol"
)

type gmsHandler struct {
	*gnet.EventServer
	codec      gnet.ICodec
	pool       *goroutine.Pool
	gnetServer gnet.Server
	// ctx    gmsContext.Context
	// cancel gmsContext.CancelFunc
	gmsServer IServer
}

func (gh *gmsHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.
	mp := protocol.MessagePack{}
	err := gh.pool.Submit(func() {
		gh.handle(mp, frame, c)
	})

	if err != nil {
		fmt.Println("[React] error:", err)
	}
	return
}

/**
处理接收到的消息
*/
func (gh *gmsHandler) handle(mp protocol.MessagePack, frame []byte, c gnet.Conn) {
	// 解析收到的二进制消息
	message, err := mp.Decode(frame)
	if err != nil {
		fmt.Println(err)
	}
	// 调用用户方法
	context, err := gh.gmsServer.HandlerMessage(message)
	if err != nil {
		fmt.Println(err)
	}
	// 获取用户方法返回的结果
	result, err := context.GetResult()
	if err != nil {
		fmt.Println(err)
	}
	// 给客户端返回处理结果
	c.AsyncWrite(result)
}

//
//
// /**
// 处理接收到的消息
// 处理粘包
// */
// func (gh *gmsHandler) handle(mp protocol.MessagePack, frame []byte, c gnet.Conn) {
//
// 	ctx := c.Context().(context.Context)
// 	connid := ctx.Value("connid").(string)
//
// 	messageCount := uint32(0)
// 	data := []byte{}
// 	for {
// 		if messageCount == 0 {
// 			data = frame
// 		} else if len(data) > int(messageCount) {
// 			fmt.Println(connid, "========11111==========")
// 			fmt.Println(connid, len(data), int(messageCount))
// 			fmt.Println(string(data))
// 			fmt.Println(connid, "==========111111========")
// 			data = data[messageCount:]
// 			fmt.Println(connid, "==========2222========")
// 			fmt.Println(string(data))
// 			fmt.Println(connid, "==========22222========")
// 		} else {
// 			break
// 		}
//
// 		message, err := mp.Decode(data)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(connid, "==============data==========")
// 		fmt.Println(string(message.GetData()))
// 		fmt.Println(connid, "==============data==========")
// 		messageCount = message.GetCount()
//
// 		context, err := gh.gmsServer.HandlerMessage(message)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		result, err := context.GetResult()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		c.AsyncWrite(result)
// 	}
//
// }

/*
gnet 服务启动成功
*/
func (gh *gmsHandler) OnInitComplete(server gnet.Server) (action gnet.Action) {
	gh.gnetServer = server
	return
}

/*
gnet 新建连接
*/
func (gh *gmsHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	// ctx, _ := gmsContext.WithCancel(gmsContext.Background())
	connid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	ctx := context.WithValue(context.Background(), "connid", connid)
	fmt.Println("[OnOpened] client: " + connid + " open." + " RemoteAddr:" + c.RemoteAddr().String())
	fmt.Println("[OnOpened] Conn count:", gh.gnetServer.CountConnections())
	c.SetContext(ctx)
	return
}

/*
gnet 连接断开
*/
func (gh *gmsHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		fmt.Println("[OnClosed] error:", err)
		return
	}
	ctx := c.Context().(context.Context)
	fmt.Println("[OnClosed] client: " + ctx.Value("connid").(string) + " Close")
	return
}
