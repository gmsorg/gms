package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type eventHandler struct {
	*gnet.EventServer
	gms  *gms
	pool *goroutine.Pool
}

func (g *eventHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("---1---")
	data := append([]byte{}, frame...)

	// Use ants pool to unblock the event-loop.
	_ = g.pool.Submit(func() {
		fmt.Println("---3---")
		// time.Sleep(1 * time.Second)
		fmt.Println(string(data))
		c.AsyncWrite(data)

	})
	fmt.Println("--2---")
	return
}

// func (g *eventHandler) React(frame []byte, c gnet.Conn) ([]byte, gnet.Action) {
// 	resMessage := &common.ResMessage{}
// 	fmt.Println(string(frame))
// 	br := common.ReqMessage{}
// 	err := json.Unmarshal(frame, &br)
// 	if err != nil {
// 		resMessage.Code = http.StatusOK
// 		resMessage.Msg = "error"
// 		b, _ := json.Marshal(resMessage)
// 		fmt.Println("return error")
// 		return b, gnet.None
// 	}
//
// 	if len(br.ServiceName) < 1 || len(br.MethodName) < 1 {
// 		return nil, gnet.None
// 	}
//
// 	// 执行方法
// 	out, err := g.exec(br.ServiceName, br.MethodName, br.ReqData)
//
// 	if err != nil {
// 		resMessage.Code = http.StatusInternalServerError
// 		resMessage.Msg = err.Error()
// 		b, _ := json.Marshal(resMessage)
// 		return b, gnet.None
// 	}
//
// 	resMessage.Code = http.StatusOK
// 	resMessage.Msg = "success"
// 	resMessage.ResData = out
//
// 	b, _ := json.Marshal(resMessage)
// 	return b, gnet.None
// }

/**
执行方法
*/
func (g *eventHandler) exec(serviceName, methodName string, reqData []byte) (res []byte, err error) {
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
