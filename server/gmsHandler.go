package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	uuid "github.com/satori/go.uuid"

	"github.com/akka/gms/igms"
)

type gmsHandler struct {
	*gnet.EventServer
	pool *goroutine.Pool
	// ctx    context.Context
	// cancel context.CancelFunc
	server igms.IServer
}

func (gh *gmsHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.
	err := gh.pool.Submit(func() {
		data := append([]byte{}, frame...)
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	if err != nil {
		// todo
	}
	return
}

/**
新建连接
*/
func (gh *gmsHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	// ctx, _ := context.WithCancel(context.Background())
	connid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	ctx := context.WithValue(context.Background(), "connid", connid)
	fmt.Println("[OnOpened] client: " + connid + " open")

	c.SetContext(ctx)
	return
}

/**
关闭连接
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
