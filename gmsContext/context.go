package gmsContext

import (
	"context"
	"encoding/json"
)

type Context struct {
	ctx        context.Context
	ReqData    []byte
	ResultData []byte
}

func NewContext() *Context {
	return &Context{
		ctx: context.Background(),
	}
}

func (c *Context) Param(param interface{}) error {
	err := json.Unmarshal(c.ReqData, param)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) Result(result interface{}) error {
	r, err := json.Marshal(result)
	if err != nil {

	}
	c.ResultData = r
	return nil
}
