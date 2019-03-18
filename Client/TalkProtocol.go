package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

const (
	PROTOCOL_INFO_TYPE_ERROR			=	0		//	信息获取错误，对应MSG为空
	PROTOCOL_INFO_TYPE_MSG				=	1		//	发送信息，对应MSG为字符串
	PROTOCOL_INFO_TYPE_START			=	2		//	启动信息，对应MSG为Patch文件信息
	PROTOCOL_INFO_TYPE_START_SUCCESS	=	3		//	启动成功，对应MSG为成功信息
	PROTOCOL_INFO_TYPE_END_SUCCESS		=	4		//	启动结束，并且成功，MSG为成功信息，可能为空
	PROTOCOL_INFO_TYPE_END_FAIL			=	5		//	启动结束，并且失败，MSG为错误信息，可能为空
	PROTOCOL_INFO_TYPE_REPORT			=	6		//	最终结果反馈，对应MSG为最终结果信息
	PROTOCOL_INFO_TYPE_OVER				=	7		//	最后信息，后续没有信息了，服务端会把链接关闭
)

type ProtocolInfo struct {
	Type		int			//	当前信息类型
	Msg			[]byte		//	当前信息内容，转成 string 就是字符串
}

func Base64Encode(Type int, Str []byte) (int, []byte) {
	var Dst []byte
	base64.StdEncoding.Encode(Dst, Str)
	if len(Dst) == 0 {
		return PROTOCOL_INFO_TYPE_ERROR, Dst
	}
	return Type, Dst
}

func Base64Decode(Type int, Str []byte) (int, []byte) {
	var Dst []byte
	l, _ := base64.StdEncoding.Decode(Dst, Str)
	if l == 0 {
		return PROTOCOL_INFO_TYPE_ERROR, Dst
	}
	return Type, Dst
}

func GetProtocolInfo(data []byte, l int) ProtocolInfo {
	v := ProtocolInfo{}
	err := json.Unmarshal(data[:l], &v)
	if err != nil {
		v.Type = PROTOCOL_INFO_TYPE_ERROR
		log.Println("GetProtocolInfo Error : ", err)
	}
	if v.Type == PROTOCOL_INFO_TYPE_MSG {
		v.Type, v.Msg = Base64Decode(v.Type, v.Msg)
		log.Println("GetProtocolInfo len = ", len(v.Msg))
	}
	return  v
}

func SetProtocolInfo(Type int, Str []byte) []byte {
	v := ProtocolInfo{}
	v.Type = Type
	v.Msg = Str
	li, _ := json.Marshal(v)
	log.Println("SetProtocolInfo len = ", len(li))
	return li
}









