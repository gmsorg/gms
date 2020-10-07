package igms

type IGms interface {
	// 启动GMS服务
	Run()
	// 注册处理器
	AddRouter(handlerName string, handlerFunc Controller)
}
