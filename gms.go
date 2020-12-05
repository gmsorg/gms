package gms

import (
	"log"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/gmsContext"
	"github.com/gmsorg/gms/plugin"
	"github.com/gmsorg/gms/server"
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

/**
默认启动参数  host+1024
*/
// func DefaultRun() {
// 	// 展示Logo
// 	common.ShowLogo()
//
// 	hostName, err := os.Hostname()
// 	if err != nil {
// 		log.Fatalf("[DefaultRun] get hostName error: %v", err)
// 	}
// 	// 启动GMS服务
// 	defaultGms.server.Run(hostName, 1024)
// }

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
