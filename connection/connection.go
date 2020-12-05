package connection

import (
	"bufio"
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gmsorg/gms/common"
	"github.com/gmsorg/gms/protocol"
)

const ReaderBuffsize = 16 * 1024

type Connection struct {
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

func (c *Connection) Send(reqData []byte) (protocol.Imessage, error) {
	wait, err := c.do(reqData)
	if err != nil {
		return nil, err
	}

	select {
	case m := <-wait:
		return m, nil
	}
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
	var err error
	for err == nil {
		mp := protocol.NewMessagePack()
		message, err := mp.ReadUnPackLen(c.reader)
		if err != nil {
			break
		}

		// fmt.Println("seq", message.GetSeq())
		//
		// fmt.Println("获取读锁")

		c.rw.RLock()
		wait, ok := c.wait[message.GetSeq()]
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
