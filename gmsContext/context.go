package gmsContext

import (
	"context"
	"encoding/json"
)

type Context struct {
	ctx        context.Context
	reqData    []byte
	resultData []byte
}

func NewContext() *Context {
	return &Context{
		ctx: context.Background(),
	}
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

	}
	c.resultData = r
	return nil
}

func (c *Context) GetResult() ([]byte, error) {
	return c.resultData, nil
}
