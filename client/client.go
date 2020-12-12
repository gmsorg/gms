package client

import (
	"errors"
	"log"
	"net"
	"sync"

	"github.com/gmsorg/gms/connection"
	"github.com/gmsorg/gms/discovery"
	"github.com/gmsorg/gms/protocol"
	"github.com/gmsorg/gms/selector"
	"github.com/gmsorg/gms/serialize"
)

type Client struct {
	rw            sync.RWMutex
	discovery     discovery.IDiscover
	selector      selector.ISelector
	connection    map[string]connection.IConnection
	messagePack   protocol.IMessagePack
	serializeType serialize.SerializeType
}

/*
NewClient 初始化客户端
*/
func NewClient(discovery discovery.IDiscover) (IClient, error) {
	client := &Client{
		discovery:     discovery,
		connection:    make(map[string]connection.IConnection),
		messagePack:   protocol.NewMessagePack(),
		serializeType: serialize.Msgpack,
	}

	client.selector = selector.NewRandomSelect(discovery)
	return client, nil
}

func (c *Client) SetSerializeType(serializeType serialize.SerializeType) error {
	if serialize := serialize.GetSerialize(c.serializeType); serialize == nil {
		return errors.New("unsupped serializeType,only supped['JSON','Msgpack','Gob']")
	}
	c.serializeType = serializeType
	return nil
}

func (c *Client) Call(serviceFunc string, request interface{}, response interface{}) error {
	connection, err := c.selector.SelectConn()
	if err != nil {
		return err
	}

	// 获取指定的序列化器
	req := serialize.GetSerialize(c.serializeType)

	// 把 request 序列化成字节数组
	serializeByte, err := req.Serialize(request)
	// fmt.Println(string(codecByte))
	if err != nil {
		log.Println(err)
	}

	// 组装消息
	message := protocol.NewMessage()
	message.SetServiceFunc(serviceFunc)
	message.SetData(serializeByte)
	message.SetSerializeType(c.serializeType)
	// todo seq
	// message.SetSeq(connection.GetSeq())
	message.SetCompressType(protocol.None)
	message.SetMessageType(protocol.Request)

	// // 打包消息
	// eb, err := c.messagePack.Pack(message, true)
	// if err != nil {
	// 	// 错误处理
	// 	log.Println(err)
	// }

	// todo 改为connection锁
	// c.rw.Lock()
	// defer c.rw.Unlock()

	// 发送打包好的消息
	// err = connection.Send(eb, response)
	err = connection.Send(message, response)

	if err != nil {
		log.Println("call-error:", err)
		if netError(err) {
			// 如果是连接错误需要清除缓存的conn对象 并清除service
			// todo 清除可用连接
			// c.cleanConn(serverKey)
		}
		return err
	}

	// // 读取返回结果消息，并解包
	// messageRes, err := connection.Read()
	// if err != nil {
	// 	return err
	// }
	//
	// res := serialize.GetSerialize(messageRes.GetSerializeType())
	// // 返序列化返回结果 成response
	// res.UnSerialize(messageRes.GetData(), response)
	return nil
}

// func (c *Client) CallOld(serviceFunc string, request interface{}, response interface{}) error {
// 	serverKey, err := c.selector.Select()
// 	if err != nil || serverKey == "" {
// 		return errors.New("can't find server")
// 	}
//
// 	// todo 锁需要优化
// 	c.rw.Lock()
// 	defer c.rw.Unlock()
//
// 	connection := c.getCachedConnection(serverKey)
//
// 	if connection == nil {
// 		// 如果是连接错误需要清除缓存的conn对象 并清除service
// 		c.cleanConn(serverKey)
// 		return errors.New("[Client.Call] generateClient fail")
// 	}
//
// 	// 获取指定的序列化器
// 	req := serialize.GetSerialize(c.serializeType)
//
// 	// 把 request 序列化成字节数组
// 	serializeByte, err := req.Serialize(request)
// 	// fmt.Println(string(codecByte))
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	// 组装消息
// 	message := protocol.NewMessage()
// 	message.SetServiceFunc(serviceFunc)
// 	message.SetData(serializeByte)
// 	message.SetSerializeType(c.serializeType)
// 	// todo seq
// 	message.SetSeq(connection.GetSeq())
// 	message.SetCompressType(protocol.None)
// 	message.SetMessageType(protocol.Request)
//
// 	eb, err := c.messagePack.Pack(message, true)
// 	if err != nil {
// 		// 错误处理
// 		log.Println(err)
// 	}
//
// 	// 发送打包好的消息
// 	messageRes, err := connection.Send(eb)
// 	if err != nil {
// 		log.Println("call-error:", err)
// 		if netError(err) {
// 			// 如果是连接错误需要清除缓存的conn对象 并清除service
// 			c.cleanConn(serverKey)
// 		}
// 		return err
// 	}
//
// 	res := serialize.GetSerialize(messageRes.GetSerializeType())
// 	// 返序列化返回结果 成response
// 	res.UnSerialize(messageRes.GetData(), response)
// 	return nil
// }

func netError(err error) bool {
	switch err.(type) {
	case net.Error:
		return true
	case *net.OpError:
		return true
	}
	return false
}

func (c *Client) getCachedConnection(address string) connection.IConnection {
	// c.rw.RLock()
	// defer c.rw.RUnlock()

	if connection, ok := c.connection[address]; ok {
		// fmt.Println("get ok", ok)
		return connection
	}
	return c.generateClient(address)
}

func (c *Client) generateClient(address string) connection.IConnection {
	newConnection := connection.NewConnection(address)
	c.connection[address] = newConnection
	return newConnection
}

func (c *Client) Close() {
	// todo
	panic("implement me")
}

/**
如果是连接错误需要清除缓存的conn对象 并清除service
*/
func (c *Client) cleanConn(serverKey string) {
	c.rw.Lock()
	defer c.rw.Unlock()

	delete(c.connection, serverKey)
	c.discovery.DeleteServer(serverKey)

}
