package client

import (
	"log"
	"net"
	"time"

	"github.com/akka/gms/common"
)

var size = 4
var width = 10
var capWidth = 16

type GmsConnection struct {
	conn    net.Conn
	bufPool *common.BytePoolCap
}

/**
创建连接
*/
func Dial(address string) (*GmsConnection, error) {
	var size = 50
	var width = 512
	var capWidth = 512

	gmsConn := &GmsConnection{
		bufPool: common.NewBytePoolCap(size, width, capWidth),
	}

	conn, err := net.DialTimeout("tcp", address, time.Second*3)
	if err != nil {
		log.Fatal(err)
	}

	gmsConn.conn = conn
	return gmsConn, nil
}
