package gms

import (
	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/gmsContext"
	"github.com/akkagao/gms/server"
)

type gms struct {
	server server.IServer
}

var defaultGms = newGms()

/*
初始化GMS
*/
func newGms() *gms {
	gms := gms{
		server: server.NewServer(),
	}
	return &gms
}

/**
添加服务路由
*/
func AddRouter(handlerName string, handlerFunc gmsContext.Controller) {
	defaultGms.server.AddRouter(handlerName, handlerFunc)
}


/*
启动GMS
*/
func Run(port int) {
	// 展示Logo
	common.ShowLogo()

	// 启动GMS服务
	defaultGms.server.Run(port)
}
