package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/akka/gms/util"
)

var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type methodType struct {
	sync.Mutex
	ReflectMethod reflect.Method
	ArgType       reflect.Type
	ReplyType     reflect.Type
}

type functionType struct {
	sync.Mutex
	fn        reflect.Value
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name         string                   // name of service
	reflectValue reflect.Value            // receiver of methods for the service
	reflectType  reflect.Type             // type of the receiver
	methods      map[string]*methodType   // registered methods
	function     map[string]*functionType // registered functions
}

/**
反射解析对象
*/
func parse(rcvr interface{}) (*service, error) {
	// 初始化 service
	service := &service{}
	service.reflectType = reflect.TypeOf(rcvr)
	service.reflectValue = reflect.ValueOf(rcvr)
	service.name = reflect.Indirect(service.reflectValue).Type().Name()
	// 是否是可导出
	if !util.IsExported(service.name) {
		return nil, errors.New(fmt.Sprintf("Register type [%v] is not exported", service.name))
	}

	// 解析方法
	service.methods = suitableMethods(service.reflectType, true)

	return service, nil
}

func suitableMethods(reflectType reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)

	for m := 0; m < reflectType.NumMethod(); m++ {
		method := reflectType.Method(m)

		mtype := method.Type
		mname := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}
		// Method needs three ins: receiver, context.Context, *args
		if mtype.NumIn() != 3 {
			if reportErr {
				log.Println("methods ", mname, " has wrong number of ins:", mtype.NumIn())
			}
			continue
		}
		// First arg must be context.Context
		ctxType := mtype.In(1)
		if !ctxType.Implements(typeOfContext) {
			if reportErr {
				log.Println("methods ", mname, " must use context.Context as the first parameter")
			}
			continue
		}

		// Second arg need not be a pointer.
		argType := mtype.In(2)
		if !isExportedOrBuiltinType(argType) {
			if reportErr {
				log.Println(mname, " parameter type not exported: ", argType)
			}
			continue
		}

		if mtype.NumOut() != 2 {
			if reportErr {
				log.Println("methods", mname, " has wrong number of outs:", mtype.NumOut())
			}
			continue
		}
		// The return type of the methods must be error.
		returnType := mtype.Out(0)
		if !isExportedOrBuiltinType(returnType) {
			if reportErr {
				log.Println(mname, " parameter type not exported: ", returnType)
			}
			continue
		}

		errType := mtype.Out(1)
		if errType != typeOfError {
			if reportErr {
				log.Println("methods", mname, " returns ", errType.String(), " not error")
			}
			continue
		}
		methods[mname] = &methodType{ReflectMethod: method, ArgType: argType, ReplyType: returnType}

		// argsReplyPools.Init(argType)
		// argsReplyPools.Init(replyType)
	}
	return methods
}

func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return util.IsExported(t.Name()) || t.PkgPath() == ""
}
