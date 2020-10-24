# GMS

今天是2020年10月24日，一年一度的程序员节，把这个还在开发中的项目开源出来，感兴趣的朋友可以一起参与开发。如果有大神路过请多指教。



GMS是一款基于[gnet](https://github.com/panjf2000/gnet)网络框架开发的Golang RPC微服务框架。

特点：

1：非常简单、学习成本极低。GMS处于初期阶段您想参与开发也非常简单。

2： 不用定义proto等协议文件。

​		写proto文件不仅麻烦还容易出错。而且使用协议文件定义服务，最终在框架内部实现都要使用反射去调用目标方法。用反射调用方法比直接调用肯定性能要好。所以如果其他条件不变的情况，用GMS这种实现方式性能肯定是最好的。

 缺点：

1：目前很多功能还没有完善，不建议应用再公司项目中。



## 快速开始

到底有多简单呢。只要您之前使用过类似Gin、beego这样的web框架，使用方法和这些web框架一样简单

下面我们以一个加法计算为类

### 1：定义请求和返回对象

```go
package model

type AdditionReq struct {
	NumberA int
	NumberB int
}

type AdditionRes struct {
	Result int
}
```

### 2：开发服务端

```go
package main

import (
	"github.com/akkagao/gms"
	"github.com/akkagao/gms/example/model"
	"github.com/akkagao/gms/gmsContext"
)

func main() {
	// 初始化GMS服务
	gms := gms.NewGms()

	// 添加业务处理路由（addition是业务处理方法的唯一标识，客户端调用需要使用）
	gms.AddRouter("addition", Addition)

	// 启动，以 1024 为启动端口
	gms.Run(1024)
}

/*
加法计算
*/
func Addition(c *gmsContext.Context) error {
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

```

### 3：开发客户端

```go
package main

import (
	"fmt"

	"github.com/akkagao/gms/client"
	"github.com/akkagao/gms/codec"
	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/example/model"
)

/*
	模拟客户端
*/
func main() {
	// 初始化一个点对点服务发现对象
	discovery := discovery.NewP2PDiscovery("127.0.0.1:1024")

	// 初始化一个客户端对象
	additionClient, err := client.NewClient(discovery)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置 Msgpack 序列化器，默认也是 Msgpack
	additionClient.SetCodecType(codec.Msgpack)

	// 请求对象
	req := &model.AdditionReq{NumberA: 10, NumberB: 20}
	// 接收返回值的对象
	res := &model.AdditionRes{}

	// 调用服务
	err = additionClient.Call("addition", req, res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%d+%d=%d", req.NumberA, req.NumberB, res.Result))
}

```



## 待开发功能

- [x] v0.1.1 服务端支持 客户端指定序列化方式 
- [ ] v0.1.2 注册中心
- [ ] 流控
- [ ] 熔断
- [ ] 监控统计

