package server

import "github.com/akka/gms/context"

type IGms interface {
	// 启动GMS服务
	Run()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc context.Controller)
}
