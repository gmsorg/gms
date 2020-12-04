package connection

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/protocol"
)

type Connection struct {
	conn        net.Conn
	messagePack protocol.IMessagePack
	seq         int64
}

func NewConnection(address string) IConnection {
	conn, err := net.DialTimeout("tcp", address, time.Second*common.ConnectTimeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Connection{
		conn:        conn,
		messagePack: protocol.NewMessagePack(),
	}
}

func (c *Connection) Send(reqData []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Connection.Send]recover send data error:%v", err)
		}
	}()
	if c.conn == nil {
		return errors.New("[Connection.Send] conn not exist")
	}
	_, err := c.conn.Write(reqData)
	return err
}

func (c *Connection) Read() (protocol.Imessage, error) {
	if c.conn == nil {
		return nil, errors.New("[Read] conn not exist")
	}

	message, err := c.messagePack.ReadUnPack(c.conn)
	if err != nil {
		return nil, fmt.Errorf("Read %w", err)
	}

	return message, nil
}
