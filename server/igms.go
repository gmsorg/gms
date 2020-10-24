package server

import "github.com/akkagao/gms/gmsContext"

type IGms interface {
	// 启动GMS服务
	Run(port int)
	// 注册处理器
	AddRouter(handlerName string, handlerFunc gmsContext.Controller)
}
