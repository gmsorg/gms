package server

import (
	"fmt"
	"sync"

	"github.com/akka/gms/igms"
)

type server struct {
	// 整个服务级别的锁
	sync.RWMutex
	// 路由Map
	routerMap map[string]igms.Controller
}

func (s *server) InitServe() {
	fmt.Println("[server] InitServe...")
}

func (s *server) Run() {
	fmt.Println("[server] run...")
}

func (s *server) Stop() {
	panic("implement me")
}

func (s *server) AddRouter(handlerName string, handlerFunc igms.Controller) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.routerMap[handlerName]; ok {
		fmt.Println("[AddRouter] fail handlerName:", handlerName, " alread exist")
		return
	}
	s.routerMap[handlerName] = handlerFunc
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
