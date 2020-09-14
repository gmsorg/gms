package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"

	"github.com/panjf2000/gnet"

	"github.com/akka/gms/common"
)

type gmsEventHandler struct {
	*gnet.EventServer
	gms *gms
}

func (g *gmsEventHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	br := common.ReqMessage{}
	err := json.Unmarshal(frame, &br)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(br.ServiceName) < 1 || len(br.MethodName) < 1 {
		return
	}

	// 执行方法
	out, err = g.exec(br.ServiceName, br.MethodName, br.ReqData)

	resMessage := &common.ResMessage{}
	if err != nil {
		resMessage.Code = http.StatusInternalServerError
		resMessage.Msg = err.Error()
		b, _ := json.Marshal(resMessage)
		return b, gnet.Action(0)
	}

	resMessage.Code = http.StatusOK
	resMessage.Msg = "success"
	resMessage.ResData = out

	b, _ := json.Marshal(resMessage)
	return b, gnet.Action(0)
}

/**
执行方法
*/
func (g *gmsEventHandler) exec(serviceName, methodName string, reqData []byte) (res []byte, err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			log.Println("recover...")
			log.Println(err1)
			err = errors.New(err1.(runtime.Error).Error())
		}
	}()

	service, exist := g.gms.serviceMap[serviceName]
	if !exist {
		return nil, errors.New(fmt.Sprintf("serviceName [%v] not exist", serviceName))
	}
	method, _ := service.methods[methodName]

	var req interface{}
	if method.ArgType.Kind() == reflect.Ptr {
		req = reflect.New(method.ArgType.Elem()).Interface()
	} else {
		req = reflect.New(method.ArgType).Interface()
	}
	err = json.Unmarshal(reqData, &req)

	// fmt.Println("=========req type=========")
	// fmt.Println(reflect.TypeOf(req))
	// fmt.Println("=========req type=========")

	if err != nil {
		log.Println("err")
		return nil, err
	}

	values := []reflect.Value{
		service.reflectValue,
		reflect.ValueOf(context.TODO()),
		reflect.ValueOf(req),
	}

	callResult := method.ReflectMethod.Func.Call(values)

	resData, err := json.Marshal(callResult[0].Interface())
	if err != nil {
		return nil, err
	}
	if callResult[1].IsNil() {
		return resData, nil
	} else {
		return reqData, callResult[1].Interface().(error)
	}
}
