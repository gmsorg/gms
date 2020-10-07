package server

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/akka/gms/example/user"
)

func TestNewServer(t *testing.T) {
	NewServer()
}

func TestGms_RegisterName(t *testing.T) {
	s := NewServer()
	Convey("TestGms_RegisterName", t, func() {
		Convey("exported", func() {
			So(s.RegisterService(new(UserServiceImpl)), ShouldBeNil)
		})
		Convey("not exported", func() {
			So(s.RegisterService(new(unUserServiceImpl)), ShouldBeError)
		})
		Convey("exit", func() {
			So(s.RegisterService(new(UserServiceImpl)), ShouldBeError)
		})
	})
}

// func TestGms_Exec(t *testing.T) {
// 	s := NewServer()
// 	jb, _ := json.Marshal(user.GetUserReq{
// 		Id: 112030,
// 	})
// 	Convey("Exec", t, func() {
// 		Convey("success", func() {
// 			So(s.RegisterService(new(UserServiceImpl)), ShouldBeNil)
// 			res, err := s.Exec("UserServiceImpl", "GetUser", jb)
// 			So(err, ShouldBeNil)
// 			So(res, ShouldNotBeNil)
// 			fmt.Println("=============")
// 			fmt.Println(res)
// 			fmt.Println("=============")
// 		})
// 	})
// }

func TestGms_Run(t *testing.T) {
	s := NewServer()
	s.Run(8080)
}

type unUserServiceImpl struct {
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) RegisterUser(c context.Context, req *user.RegisterUserReq) (*user.RegisterUserRes, error) {
	fmt.Println(req.Name, req.Barthday, req.Sex)
	res := &user.RegisterUserRes{}
	res.Id = rand.Int63n(1000000)
	return res, nil
}

func (u *UserServiceImpl) GetUser(c context.Context, req *user.GetUserReq) (*user.GetUserRes, error) {
	fmt.Println(fmt.Sprintf("获取ID：%v 的用户", req.Id))
	res := &user.GetUserRes{}

	// i := 0
	// fmt.Println(1 / i)
	res.Name = "CrazyWolf"
	res.Sex = "nan"
	res.Barthday = "2020-09-02"
	return res, nil
}

