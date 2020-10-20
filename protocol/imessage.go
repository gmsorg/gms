package protocol

import "github.com/akkagao/gms/codec"

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
	// 设置扩展数据的长度
	SetExtLen(extLen uint32)
	// 获取扩展数据的长度
	GetExtLen() uint32
	// 设置扩展数据
	SetExt(ext []byte)
	// 获取扩展数据
	GetExt() []byte

	// 设置编码方式
	SetCodecType(codecType codec.CodecType)
	// 获取编码方式
	GetCodecType() codec.CodecType

	// 设置主体数据段长度
	SetDataLen(dataLen uint32)
	// 获取主体数据段长度
	GetDataLen() uint32
	// 设置主体数据内容
	SetData(data []byte)
	// 获取主体数据内容
	GetData() []byte

	// 获取消息总长度
	GetCount() uint32
}
