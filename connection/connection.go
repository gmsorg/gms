package connection

import (
	"bufio"
	"errors"
	"log"
	"net"
	"sync"
	"sync/atomic"
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
	waitM       common.Map

	// rw    sync.RWMutex
}
type callCmd struct {
	start int64
	cost  time.Duration
	// sess           *session
	// output         Message
	result interface{}
	// stat           *Status
	// inputMeta      *utils.Args
	// swap           goutil.Map
	mu             sync.Mutex
	callCmdChan    chan<- callCmd // Send itself to the public channel when call is complete.
	doneChan       chan struct{}  // Strobes when call is complete.
	inputBodyCodec byte
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
		// wait:        make(map[uint64]chan protocol.Imessage, 10),
		waitM: common.NewAtomicMap(),
	}

	c.reader = bufio.NewReaderSize(conn, ReaderBuffsize)
	go c.read()

	return c
}

func (c *Connection) SetConnId(connId int) {
	c.connId = connId
}

func (c *Connection) Send(message protocol.Imessage, response interface{}) error {

	if c.conn == nil {
		return errors.New("[Connection.Send] conn not exist")
	}

	wait, err := c.AsyncSend(message, response)

	if err != nil {
		return err
	}

	<-wait.doneChan

	return nil
}

func (c *Connection) AsyncSend(message protocol.Imessage, response interface{}) (*callCmd, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Connection.Send]recover send data error:%v", err)
		}
	}()

	cc := &callCmd{
		start:  0,
		cost:   0,
		result: response,
		// callCmdChan:    nil,
		doneChan: make(chan struct{}),
		// inputBodyCodec: 0,
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()

	seq := atomic.AddUint64(&c.seq, 1)
	message.SetSeq(seq)

	c.waitM.Store(seq, cc)
	// c.rw.Unlock()

	// 打包消息
	eb, err := c.messagePack.Pack(message, true)
	if err != nil {
		// c.rw.Lock()
		// delete(c.waitM, seq)
		c.waitM.Delete(seq)
		// c.rw.Unlock()
		// 错误处理
		log.Println(err)
	}

	// c.rw.Lock()
	// defer c.rw.Unlock()
	_, err = c.conn.Write(eb)
	if err != nil {
		return nil, err
	}
	return cc, nil
}

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

		seq := message.GetSeq()
		wait, ok := c.waitM.Load(seq)
		c.waitM.Delete(seq)

		if ok && wait != nil {
			switch w := wait.(type) {
			case *callCmd:

				serialize := serialize.GetSerialize(message.GetSerializeType())
				// 返序列化返回结果 成response
				if err := serialize.UnSerialize(message.GetData(), w.result); err != nil {
					log.Println(err)
					break
				}

				w.doneChan <- struct{}{}

			}

		}
	}
}
