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

func (c *Context) Param(param interface{}) error {
	err := json.Unmarshal(c.reqData, param)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) Result(result interface{}) error {
	r, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(r))
	c.resultData = r
	return nil
}

func (c *Context) GetResult() ([]byte, error) {
	return c.resultData, nil
}
