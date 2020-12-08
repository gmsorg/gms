package connection

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/protocol"
	"github.com/gmsorg/gms/serialize"
)

const ReaderBuffsize = 16 * 1024

type Connection struct {
	connId      int
	conn        net.Conn
	messagePack protocol.IMessagePack
	seq         uint64
	reader      *bufio.Reader
	wait        map[uint64]chan protocol.Imessage
	rw          sync.RWMutex
}

// type Wait struct {
// 	finish chan protocol.Imessage
// }

func NewConnection(address string) IConnection {
	conn, err := net.DialTimeout("tcp", address, time.Second*common.ConnectTimeout)
	if err != nil {
		log.Println(err)
		return nil
	}

	c := &Connection{
		conn:        conn,
		messagePack: protocol.NewMessagePack(),
		wait:        make(map[uint64]chan protocol.Imessage, 10),
	}

	c.reader = bufio.NewReaderSize(conn, ReaderBuffsize)
	go c.read()

	return c
}

func (c *Connection) SetConnId(connId int) {
	c.connId = connId
}

func (c *Connection) SendM(message protocol.Imessage, response interface{}) error {

	if c.conn == nil {
		return errors.New("[Connection.Send] conn not exist")
	}

	wait, err := c.doM(message)

	if err != nil {
		return err
	}

	select {
	case m := <-wait:
		res := serialize.GetSerialize(m.GetSerializeType())
		// 返序列化返回结果 成response
		res.UnSerialize(m.GetData(), response)
	}
	return nil
}

func (c *Connection) doM(message protocol.Imessage) (chan protocol.Imessage, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Connection.Send]recover send data error:%v", err)
		}
	}()

	c.rw.Lock()
	defer c.rw.Unlock()

	seq := c.seq
	c.seq++
	message.SetSeq(seq)

	wait := make(chan protocol.Imessage)
	// fmt.Println("获取写锁")

	// fmt.Println("获取写锁成功")
	c.wait[seq] = wait
	// fmt.Println("set seq:", seq)
	// fmt.Println("释放写锁成功")

	// 打包消息
	eb, err := c.messagePack.Pack(message, true)
	if err != nil {
		delete(c.wait, seq)
		// 错误处理
		log.Println(err)
	}

	_, err = c.conn.Write(eb)
	if err != nil {
		return nil, err
	}
	return wait, nil
}

func (c *Connection) Send(reqData []byte, response interface{}) error {
	// c.rw.Lock()
	// defer c.rw.Unlock()

	wait, err := c.do(reqData)
	if err != nil {
		return err
	}

	select {
	case m := <-wait:
		res := serialize.GetSerialize(m.GetSerializeType())
		// 返序列化返回结果 成response
		res.UnSerialize(m.GetData(), response)
	}
	return nil
}

func (c *Connection) GetSeq() uint64 {
	return c.seq
}

func (c *Connection) do(reqData []byte) (chan protocol.Imessage, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Connection.Send]recover send data error:%v", err)
		}
	}()

	seq := c.seq
	c.seq++
	wait := make(chan protocol.Imessage)

	// fmt.Println("获取写锁")
	c.rw.Lock()
	// fmt.Println("获取写锁成功")
	c.wait[seq] = wait
	// fmt.Println("set seq:", seq)
	c.rw.Unlock()
	// fmt.Println("释放写锁成功")

	if c.conn == nil {
		return nil, errors.New("[Connection.Send] conn not exist")
	}
	_, err := c.conn.Write(reqData)
	if err != nil {
		return nil, err
	}
	return wait, nil
}

// func (c *Connection) Read() (protocol.Imessage, error) {
// 	if c.conn == nil {
// 		return nil, errors.New("[Read] conn not exist")
// 	}
//
// 	message, err := c.messagePack.ReadUnPackLen(c.conn)
// 	if err != nil {
// 		return nil, fmt.Errorf("Read %w", err)
// 	}
//
// 	return message, nil
// }

/**
读取返回信息
*/
func (c *Connection) read() {
	fmt.Println("connId:", c.connId)
	var err error
	for err == nil {
		mp := protocol.NewMessagePack()
		message, err := mp.ReadUnPackLen(c.reader)
		if err != nil {
			break
		}

		// fmt.Println("get seq:", message.GetSeq())
		//
		// fmt.Println("获取读锁")

		c.rw.RLock()
		wait, ok := c.wait[message.GetSeq()]
		delete(c.wait, message.GetSeq())
		c.rw.RUnlock()

		// fmt.Println("获取读锁2")

		if ok {
			if wait == nil {
				break
			}
			wait <- message
		}
	}
}
