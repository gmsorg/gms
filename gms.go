package gms

import (
	"github.com/akka/gms/common"
	"github.com/akka/gms/igms"
	"github.com/akka/gms/server"
)

type gms struct {
	server igms.IServer
}

/**
初始化GMS
*/
func NewGms() igms.IGms {
	gms := gms{
		server: server.NewServer(),
	}
	return &gms
}

/**
添加服务路由
*/
func (g *gms) AddRouter(handlerName string, handlerFunc igms.Controller) {
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
