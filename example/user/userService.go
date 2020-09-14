package user

import (
	"context"
)

type UserService interface {
	// 注册用户
	RegisterUser(c context.Context, req *RegisterUserReq) (*RegisterUserRes, error)

	// 查询用户
	GetUser(c context.Context, req *GetUserReq) (*GetUserRes, error)
}
