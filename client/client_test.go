package client

import (
	"context"
	"log"
	"testing"

	"github.com/akka/gms/example/user"
)

func Test_gmsClient_Call(t *testing.T) {
	conn, err := Dial("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	NewUserServiceClient(conn)

	conn.Call()
}

type UserServiceClient struct {
}

func NewUserServiceClient(conn *gmsConnection) (*UserServiceClient, error) {
	return nil, nil
}

func (u *UserServiceClient) RegisterUser(c context.Context, req *user.RegisterUserReq) (*user.RegisterUserRes, error) {
	panic("implement me")
}

func (u *UserServiceClient) GetUser(c context.Context, req *user.GetUserReq) (*user.GetUserRes, error) {
	panic("implement me")
}
