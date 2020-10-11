package gmsContext

import (
	"context"
	"encoding/json"
	"fmt"
)

type Context struct {
	context.Context // todo context 功能待完善（参考gin的context实现）
	reqData         []byte
	resultData      []byte
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetParam(b []byte) {
	c.reqData = b
}

/**
把请求中的信息反序列化成用户指定的对象
*/
func (c *Context) Param(param interface{}) error {
	// todo 改为其他序列化方式
	// fmt.Println("===========")
	// fmt.Println(string(c.reqData))
	// fmt.Println("===========")
	err := json.Unmarshal(c.reqData, param)
	if err != nil {
		fmt.Println("[Param] error", err)
		return err
	}
	return nil
}

func (c *Context) Result(result interface{}) error {
	// todo 改为其他序列化方式
	r, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(len(r))
	c.resultData = r
	return nil
}

func (c *Context) GetResult() ([]byte, error) {
	return c.resultData, nil
}
