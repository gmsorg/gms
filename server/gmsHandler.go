package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/satori/go.uuid"
)

type gmsHandler struct {
	*gnet.EventServer
	pool       *goroutine.Pool
	gnetServer gnet.Server
	// ctx    context.Context
	// cancel context.CancelFunc
	gmsServer IServer
}

func (gh *gmsHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.
	err := gh.pool.Submit(func() {
		// data := append([]byte{}, frame...)
		// time.Sleep(1 * time.Second)
		c.AsyncWrite(frame)
	})

	if err != nil {
		fmt.Println("[React] error:", err)
	}
	return
}

/**
gnet 服务启动成功
*/
func (gh *gmsHandler) OnInitComplete(server gnet.Server) (action gnet.Action) {
	gh.gnetServer = server
	return
}

/**
gnet 新建连接
*/
func (gh *gmsHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	// ctx, _ := context.WithCancel(context.Background())
	connid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	ctx := context.WithValue(context.Background(), "connid", connid)
	fmt.Println("[OnOpened] client: " + connid + " open." + " RemoteAddr:" + c.RemoteAddr().String())
	fmt.Println("[OnOpened] Conn count:", gh.gnetServer.CountConnections())
	c.SetContext(ctx)
	return
}

/**
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
