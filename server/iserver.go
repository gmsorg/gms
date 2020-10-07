package server

import "github.com/akka/gms/context"

type IServer interface {
	// 初始化GMS服务
	InitServe()
	// 启动GMS服务
	Run()
	// 停止GMS服务
	Stop()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc context.Controller)
}
