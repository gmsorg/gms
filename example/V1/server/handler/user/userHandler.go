package user

import (
	"github.com/akkagao/gms/common"
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

	// fmt.Println(addUserReq)

	res := &vo.AddUserRes{}
	res.OrgName = addUserReq.Name
	res.NewName = "hahahha" + common.GenIdentity()

	c.Result(res)
	return nil
}
