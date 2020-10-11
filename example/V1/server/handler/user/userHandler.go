package user

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/akkagao/gms/example/V1/vo"
	"github.com/akkagao/gms/gmsContext"
)

/*
UserAdd 测试方法
*/
func UserAdd(c *gmsContext.Context) error {
	// fmt.Println("call userAdd...")
	addUserReq := &vo.AddUserReq{}
	c.Param(addUserReq)

	fmt.Println(addUserReq)

	addUserReq.OrgName = addUserReq.Name
	addUserReq.Name = "hahahha" + uuid.NewV4().String()

	c.Result(addUserReq)
	return nil
}
