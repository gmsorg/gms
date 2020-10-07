package context

import "encoding/json"

type Context struct {
	Data []byte
}

func (c *Context) Param(param interface{}) error {
	err := json.Unmarshal(c.Data, param)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) Result(result interface{}) error {
	panic("implement me")
}
