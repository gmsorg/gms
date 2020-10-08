package user

import (
	"fmt"

	"github.com/akka/gms/example/V1/vo"
	"github.com/akka/gms/gmsContext"
)

func UserAdd(c *gmsContext.Context) error {
	fmt.Println("call userAdd...")
	addUserReq := &vo.AddUserReq{}
	c.Param(addUserReq)

	fmt.Println(addUserReq)
	return nil
}
