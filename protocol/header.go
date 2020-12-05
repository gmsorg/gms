package protocol

import "github.com/akkagao/gms/common"

const (
	magicNumber byte = 0x88
)

func MagicNumber() byte {
	return magicNumber
}

// 消息类型
type MessageType byte

const (
	// 心跳消息
	Heartbeat MessageType = iota
	// Request 消息
	Request
	// 正常返回消息
	Response
	// 错误返回消息
	ResponseError
)

type CompressType byte

const (
	None CompressType = iota
	Gzip
)

type Header [common.HeaderLength]byte
