package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"

	"github.com/akka/gms/common"
	"github.com/akka/gms/context"
)

type server struct {
	// 整个服务级别的锁
	sync.RWMutex
	// 路由Map
	routerMap map[string]context.Controller
	// gms 服务
	gmsHandler *gmsHandler
}

/**
初始化GMS服务
*/
func NewServer() IServer {
	s := server{
		routerMap: make(map[string]context.Controller),
	}
	return &s
}

/**
准备启动服务的资源
*/
func (s *server) InitServe() {
	fmt.Println("[gmsServer] InitServe")

	pool := goroutine.Default()
	defer pool.Release()

	codec = gnet.NewLengthFieldBasedFrameCodec(encoderConfig, decoderConfig)
	// 启动gnet
	s.gmsHandler = &gmsHandler{
		gmsServer: s,
		pool:      pool,
	}
	log.Fatal(gnet.Serve(s.gmsHandler, fmt.Sprintf("tcp://:%v", common.GMS_PORT), gnet.WithMulticore(true)))
}

/**
启动服务
*/
func (s *server) Run() {
	fmt.Println("[gmsServer] start run gms gmsServer")
	// 准备启动服务的资源
	s.InitServe()

}

/**
停止服务 回收资源
*/
func (s *server) Stop() {
	fmt.Println("[gmsServer] stop")
}

/**
添加路由
*/
func (s *server) AddRouter(handlerName string, handlerFunc context.Controller) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.routerMap[handlerName]; ok {
		fmt.Println("[AddRouter] fail handlerName:", handlerName, " alread exist")
		return
	}
	s.routerMap[handlerName] = handlerFunc
}
