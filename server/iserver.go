package server

import (
	"github.com/akkagao/gms/gmsContext"
	"github.com/akkagao/gms/protocol"
)

type IServer interface {
	// 初始化GMS服务
	InitServe()
	// 启动GMS服务
	Run()
	// 停止GMS服务
	Stop()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc gmsContext.Controller)
	// 获取处理器
	GetRouter(handlerName string) (gmsContext.Controller, error)
	// 处理消息
	HandlerMessage(message protocol.Imessage) (*gmsContext.Context, error)
}
