package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/akka/gms/example/user"
)

func Test_gmsClient_Call_GO(t *testing.T) {
	s := time.Now()
	conn, err := Dial("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	userServiceClient, err := NewUserServiceClient(conn)

	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		for i := 0; i < 2; i++ {
			go func() {
				begin := time.Now()
				getUserRes, _ := userServiceClient.RegisterUser(context.Background(), &user.RegisterUserReq{
					Name: uuid.NewV4().String(),
				})

				fmt.Println(getUserRes)
				fmt.Println(time.Since(begin))
				wg.Done()
			}()

		}
	}
	wg.Wait()

	fmt.Println("==============")
	fmt.Println(time.Since(s))
}

func Test_gmsClient_Call(t *testing.T) {
	s := time.Now()
	conn, err := Dial("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	userServiceClient, err := NewUserServiceClient(conn)

	for i := 0; i < 50; i++ {

		begin := time.Now()
		getUserRes, _ := userServiceClient.RegisterUser(context.Background(), &user.RegisterUserReq{})

		fmt.Println(getUserRes)
		fmt.Println(time.Since(begin))
	}

	fmt.Println("==============")
	fmt.Println(time.Since(s))

	// conn.Call()
}

type UserServiceClient struct {
	conn *GmsConnection
}

func NewUserServiceClient(conn *GmsConnection) (*UserServiceClient, error) {
	serviceClient := &UserServiceClient{
		conn: conn,
	}
	return serviceClient, nil
}

func (u *UserServiceClient) RegisterUser(c context.Context, req *user.RegisterUserReq) (*user.RegisterUserRes, error) {
	serviceName := "UserServiceImpl"
	methodName := "RegisterUser"

	res := &user.RegisterUserRes{}

	u.conn.CommonCall(serviceName, methodName, req, res)

	return res, nil
}

func (u *UserServiceClient) GetUser(c context.Context, req *user.GetUserReq) (*user.GetUserRes, error) {
	return nil, nil
}
