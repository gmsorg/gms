package gms

import (
	"log"

	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/gmsContext"
	"github.com/akkagao/gms/plugin"
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

/**
注册插件
*/
func AddPlugin(plugin plugin.IPlugin) {
	defaultGms.server.AddPlugin(plugin)
}

/*
启动GMS
*/
func Run(ip string, port int) {
	// 校验IP是否正确
	err := common.ValidateIp(ip)
	if err != nil {
		log.Fatalf("ip: %v error: %v", ip, err)
	}

	// 校验端口是否正确
	err = common.ValidatePort(port)
	if err != nil {
		log.Fatalf("port: %v error: %v", port, err)
	}

	// 展示Logo
	common.ShowLogo()

	// 启动GMS服务
	defaultGms.server.Run(ip, port)
}
