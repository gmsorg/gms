package gms

import (
	"github.com/akka/gms/common"
	"github.com/akka/gms/gmsContext"
	"github.com/akka/gms/server"
)

type gms struct {
	server server.IServer
}

/**
初始化GMS
*/
func NewGms() server.IGms {
	gms := gms{
		server: server.NewServer(),
	}
	return &gms
}

/**
添加服务路由
*/
func (g *gms) AddRouter(handlerName string, handlerFunc gmsContext.Controller) {
	g.server.AddRouter(handlerName, handlerFunc)
}

/**
启动GMS
*/
func (g *gms) Run() {
	// 展示Logo
	common.ShowLogo()

	// 启动GMS服务
	g.server.Run()
}
