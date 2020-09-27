package service

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet"

	"github.com/akka/gms/common"
)

type gms struct {
	sync.Mutex
	serviceMap map[string]*service
}

/**
初始化Gms
*/
func NewServer() *gms {
	common.Show_logo()

	server := &gms{
		serviceMap: map[string]*service{},
	}
	return server
}

/**
注册服务
*/
func (gms *gms) RegisterService(rcvr interface{}) error {
	gms.Lock()
	defer gms.Unlock()

	service, err := parse(rcvr)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, ok := gms.serviceMap[service.name]; ok {
		return errors.New(fmt.Sprintf("Register type [%v] alread exist", service.name))
	}

	gms.serviceMap[service.name] = service

	return nil
}

/**
启动GMS
*/
func (gms *gms) Run(port int) error {
	// 初始化时间处理器
	gmsEventHandler := &eventHandler{gms: gms}

	log.Println("GMS service listen on:", port)

	// 启动gnet，设置由 eventHandler 处理请求信息
	err := gnet.Serve(gmsEventHandler,
		fmt.Sprintf("tcp://:%d", port),
		gnet.WithMulticore(true),
		gnet.WithReusePort(true))

	return err
}
