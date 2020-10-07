package server

import (
	"time"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"

	"github.com/akka/gms/igms"
)

type gmsHandler struct {
	*gnet.EventServer
	server igms.IServer
	pool   *goroutine.Pool
}

func (es *gmsHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.
	_ = es.pool.Submit(func() {
		data := append([]byte{}, frame...)
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	return
}
