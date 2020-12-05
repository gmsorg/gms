package server

import (
	"github.com/gmsorg/gms/gmsContext"
	"github.com/gmsorg/gms/plugin"
	"github.com/gmsorg/gms/protocol"
)

type IServer interface {
	// 初始化GMS服务
	InitServe(port int)
	// 启动GMS服务
	Run(ip string, port int)
	// 停止GMS服务
	Stop()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc gmsContext.Controller)
	// 获取处理器
	GetRouter(handlerName string) (gmsContext.Controller, error)
	// 处理消息
	HandlerMessage(message protocol.Imessage) (*gmsContext.Context, error)
	// 注册插件
	AddPlugin(plugin plugin.IPlugin)
}
