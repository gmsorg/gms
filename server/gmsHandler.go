package server

import (
	"context"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/protocol"
)

type gmsHandler struct {
	*gnet.EventServer
	codec      gnet.ICodec
	pool       *goroutine.Pool
	gnetServer gnet.Server
	// ctx    gmsContext.Context
	// cancel gmsContext.CancelFunc
	messagePack protocol.IMessagePack
	gmsServer   IServer
}

func (gh *gmsHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.

	err := gh.pool.Submit(func() {
		gh.handle(frame, c)
	})

	if err != nil {
		log.Println("[React] error:", err)
	}
	return
}

/**
处理接收到的消息
*/
func (gh *gmsHandler) handle(frame []byte, c gnet.Conn) {
	// 解析收到的二进制消息
	message, err := gh.messagePack.UnPack(frame)
	if err != nil {
		log.Println(err)
	}
	// 调用用户方法
	context, err := gh.gmsServer.HandlerMessage(message)
	if err != nil {
		log.Println(err)
	}
	// 获取用户方法返回的结果
	result, err := context.GetResult()
	if err != nil {
		log.Println(err)
	}

	resultMessage := protocol.NewMessage()
	resultMessage.SetData(result)
	resultMessage.SetSerializeType(message.GetSerializeType())
	resultMessage.SetSeq(message.GetSeq())
	resultMessage.SetCompressType(message.GetCompressType())
	resultMessage.SetMessageType(protocol.Response)

	rb, err := gh.messagePack.Pack(resultMessage, false)
	if err != nil {
		log.Println("[gmsHandler handle] error: %v", err)
	}
	// 给客户端返回处理结果
	c.AsyncWrite(rb)
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
// 			log.Println(connid, "========11111==========")
// 			log.Println(connid, len(data), int(messageCount))
// 			log.Println(string(data))
// 			log.Println(connid, "==========111111========")
// 			data = data[messageCount:]
// 			log.Println(connid, "==========2222========")
// 			log.Println(string(data))
// 			log.Println(connid, "==========22222========")
// 		} else {
// 			break
// 		}
//
// 		message, err := mp.UnPack(data)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		log.Println(connid, "==============data==========")
// 		log.Println(string(message.GetData()))
// 		log.Println(connid, "==============data==========")
// 		messageCount = message.GetCount()
//
// 		context, err := gh.gmsServer.HandlerMessage(message)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		result, err := context.GetResult()
// 		if err != nil {
// 			log.Println(err)
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
	connid := common.GenIdentity()
	ctx := context.WithValue(context.Background(), "connid", connid)
	log.Println(fmt.Sprintf("[OnOpened] client: %v open. RemoteAddr:%v", connid, c.RemoteAddr().String()))
	log.Println("[OnOpened] Conn count:", gh.gnetServer.CountConnections())
	c.SetContext(ctx)
	return
}

/*
gnet 连接断开
*/
func (gh *gmsHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Println("[OnClosed] error:", err)
		return
	}
	ctx := c.Context().(context.Context)
	log.Println("[OnClosed] client: " + ctx.Value("connid").(string) + " Close")
	return
}
