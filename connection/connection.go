package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/akkagao/gms/common"
	"github.com/akkagao/gms/protocol"
)

type Connection struct {
	conn        net.Conn
	messagePack protocol.IMessagePack
}

func NewConnection(address string) IConnection {
	conn, err := net.DialTimeout("tcp", address, time.Second*common.ConnectTimeout)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &Connection{
		conn:        conn,
		messagePack: protocol.NewMessagePack(),
	}
}

func (c *Connection) Send(reqData []byte) error {
	if c.conn == nil {
		return errors.New("[Send] conn not exist")
	}
	_, err := c.conn.Write(reqData)
	return err
}

func (c *Connection) Read(response interface{}) error {
	if c.conn == nil {
		return errors.New("[Read] conn not exist")
	}

	message, err := c.messagePack.ReadUnPack(c.conn)
	if err != nil {
		return fmt.Errorf("Read %v", err)
	}

	json.Unmarshal(message.GetData(), response)

	return nil
}
