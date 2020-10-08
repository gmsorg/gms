package server

import "github.com/akka/gms/gmsContext"

type IGms interface {
	// 启动GMS服务
	Run()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc gmsContext.Controller)
}
