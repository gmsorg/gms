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

	bytes := common.BytePool.Get()
	defer common.BytePool.Put(bytes)

	// len := 0
	//
	// for {
	// 	n, err := c.conn.Read(bytes[len:])
	// 	if n > 0 {
	// 		len += n
	// 	}
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			// Error Handler
	// 		}
	//
	// 		break
	// 	}
	// }

	return nil
}
