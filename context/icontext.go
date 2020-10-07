package context

type Controller func(c *Context) error

type IContext interface {
	// 参数转换
	Param(interface{}) error
	// 返回参数
	Result(interface{}) error
}
