package user

import (
	"fmt"

	"github.com/akka/gms/context"
)

type AddUserReq struct {
	Name string
}

func UserAdd(c *context.Context) error {
	fmt.Println("call userAdd...")
	addUserReq := &AddUserReq{}
	c.Param(addUserReq)

	fmt.Println(addUserReq)
	return nil
}
