package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"

	"github.com/akka/gms/common"
	"github.com/akka/gms/igms"
)

type server struct {
	// 整个服务级别的锁
	sync.RWMutex
	// 路由Map
	routerMap map[string]igms.Controller
}

/**
初始化GMS服务
*/
func NewServer() igms.IServer {
	s := server{
		routerMap: make(map[string]igms.Controller),
	}
	return &s
}

/**
准备启动服务的资源
*/
func (s *server) InitServe() {
	fmt.Println("[server] InitServe")

	pool := goroutine.Default()
	defer pool.Release()

	// 启动gnet
	gmsHandler := &gmsHandler{
		server: s,
		pool:   pool,
	}
	log.Fatal(gnet.Serve(gmsHandler, fmt.Sprintf("tcp://:%v", common.GMS_PORT), gnet.WithMulticore(true)))
}

/**
启动服务
*/
func (s *server) Run() {
	fmt.Println("[server] start run gms server")
	// 准备启动服务的资源
	s.InitServe()

}

/**
停止服务 回收资源
*/
func (s *server) Stop() {
	fmt.Println("[server] stop")
}

/**
添加路由
*/
func (s *server) AddRouter(handlerName string, handlerFunc igms.Controller) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.routerMap[handlerName]; ok {
		fmt.Println("[AddRouter] fail handlerName:", handlerName, " alread exist")
		return
	}
	s.routerMap[handlerName] = handlerFunc
}
