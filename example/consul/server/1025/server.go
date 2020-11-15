package main

import (
	"log"

	"github.com/akkagao/gms"
	"github.com/akkagao/gms/example/model"
	"github.com/akkagao/gms/plugin/registe"

	"github.com/akkagao/gms/gmsContext"
)

func main() {
	// 添加业务处理路由（addition是业务处理方法的唯一标识，客户端调用需要使用）
	gms.AddRouter("addition", Addition)

	// 注册redis服务发现插件
	gms.AddPlugin(registe.NewConsulRegistePlugin("gmsDemo", []string{"127.0.0.1:8500"}))

	// 启动，以1024 为启动端口
	gms.Run("127.0.0.1", 1025)
	// gms.DefaultRun()
}

/*
加法计算
*/
func Addition(c *gmsContext.Context) error {
	log.Println(1025)
	additionReq := &model.AdditionReq{}
	// 绑定请求参数
	c.Param(additionReq)

	// 结果对象
	additionRes := &model.AdditionRes{}
	additionRes.Result = additionReq.NumberA + additionReq.NumberB

	// 返回结果
	c.Result(additionRes)
	return nil
}
