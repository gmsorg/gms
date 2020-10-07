package igms

import "github.com/akka/gms/context"

type Controller func(c *context.Context) error

type IContext interface {
	// 参数转换
	Param(interface{}) error
	// 返回参数
	Result(interface{}) error
}
