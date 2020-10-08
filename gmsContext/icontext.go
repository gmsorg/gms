package gmsContext

type Controller func(c *Context) error

type IContext interface {
	// 设置参数数据
	SetParam(b []byte) error
	// 参数转换
	Param(interface{}) error
	// 返回参数
	Result(interface{}) error
	// 获取结果
	GetResult() ([]byte, error)
}
