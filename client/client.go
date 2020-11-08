package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/akkagao/gms/codec"
	"github.com/akkagao/gms/connection"
	"github.com/akkagao/gms/discovery"
	"github.com/akkagao/gms/protocol"
	"github.com/akkagao/gms/selector"
)

type Client struct {
	rw          sync.RWMutex
	discovery   discovery.IDiscover
	selector    selector.ISelector
	connection  map[string]connection.IConnection
	messagePack protocol.IMessagePack
	codecType   codec.CodecType
}

/*
NewClient 初始化客户端
*/
func NewClient(discovery discovery.IDiscover) (IClient, error) {
	client := &Client{
		discovery:   discovery,
		connection:  make(map[string]connection.IConnection),
		messagePack: protocol.NewMessagePack(),
		codecType:   codec.Msgpack,
	}

	client.selector = selector.NewRandomSelect(discovery)
	return client, nil
}

func (c *Client) SetCodecType(codecType codec.CodecType) error {
	if codec := codec.GetCodec(c.codecType); codec == nil {
		return errors.New("unsupped codecType,only supped['JSON','Msgpack','Gob']")
	}
	c.codecType = codecType
	return nil
}

func (c *Client) Call(serviceFunc string, request interface{}, response interface{}) error {
	serverKey, err := c.selector.Select()
	if err != nil || serverKey == "" {
		return errors.New("can't find server")
	}

	connection := c.getCachedConnection(serverKey)
	if connection == nil {
		connection = c.generateClient(serverKey)
	}
	if connection == nil {
		// 如果是连接错误需要清除缓存的conn对象 并清除service
		c.cleanConn(serverKey)
		return errors.New("[Client.Call] generateClient fail")
	}

	// 获取指定的序列化器
	codecReq := codec.GetCodec(c.codecType)

	// 把 request 序列化成字节数组
	codecByte, err := codecReq.Encode(request)
	if err != nil {
		log.Println(err)
	}

	// 组装消息
	message := protocol.NewMessage([]byte(serviceFunc), codecByte, c.codecType)

	// 打包消息
	eb, err := c.messagePack.Pack(message)
	if err != nil {
		// 错误处理
		log.Println(err)
	}

	// 发送打包好的消息
	err = connection.Send(eb)
	if err != nil {
		fmt.Println("call-error:", err)
		if netError(err) {
			// 如果是连接错误需要清除缓存的conn对象 并清除service
			c.cleanConn(serverKey)
		}
		return err
	}

	// 读取返回结果消息，并解包
	messageRes, err := connection.Read()
	if err != nil {
		return err
	}

	codecRes := codec.GetCodec(messageRes.GetCodecType())
	// 返序列化返回结果 成response
	codecRes.Decode(messageRes.GetData(), response)
	return nil
}

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
	c.rw.RLock()
	defer c.rw.RUnlock()
	if connection, ok := c.connection[address]; ok {
		return connection
	}
	return nil
}

func (c *Client) generateClient(address string) connection.IConnection {
	c.rw.Lock()
	defer c.rw.Unlock()

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
