package connection

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/akkagao/gms/common"
)

type Connection struct {
	conn net.Conn
}

func (c *Connection) Send(data []byte) error {
	if c.conn == nil {
		return errors.New("conn not exist")
	}
	_, err := c.conn.Write(data)
	return err
}

func NewConnection(address string) IConnection {
	conn, err := net.DialTimeout("tcp", address, time.Second*common.ConnectTimeout)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &Connection{
		conn: conn,
	}
}
