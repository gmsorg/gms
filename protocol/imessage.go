package protocol

import (
	"github.com/akkagao/gms/codec"
)

/*
Imessage 请求消息和返回消息体封装

消息格式
扩展数据长度|主体数据长度|编码方式|扩展数据|主体数据
扩展数据可以按照实际使用场景定义格式和用途

在gms服务中
请求消息 扩展信息作为：要请求的目标方法
返回消息 扩展信息作为：请求成功失败的描述
*/
type Imessage interface {
	GetHeader() Header

	CheckMagicNumber() bool

	Version() byte

	SetVersion(v byte)

	MessageType() MessageType

	SetMessageType(mt MessageType)

	// 获取消息压缩类型
	CompressType() CompressType

	// 设置消息压缩类型
	SetCompressType(ct CompressType)

	// 获取序列化类型
	SerializeType() codec.CodecType
	// 设置序列化类型
	SetSerializeType(ct codec.CodecType)

	// 获取消息序号
	Seq() uint64
	// 设置消息序号
	SetSeq(seq uint64)

	// 获取请求方法名
	ServiceFunc() string
	// 设置请求方法名
	SetServiceFunc(serviceFunc string)

	// 设置扩展数据
	SetExt(ext map[string]string)
	// 获取扩展数据
	GetExt() map[string]string

	// 设置主体数据内容
	SetData(data []byte)
	// 获取主体数据内容
	GetData() []byte
}
