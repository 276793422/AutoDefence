package main

import (
	"encoding/base64"
	"encoding/json"
)

//定义配置文件解析后的结构
type BodyMsg struct {
	Type	string
	Msg		string
}

func Base64Encode(Str []byte) (int, []byte) {
	var Dst []byte
	base64.StdEncoding.Encode(Dst, Str)
	if len(Dst) == 0 {
		return 0, Dst
	}
	return len(Dst), Dst
}

func Base64Decode(Str string) (int, string) {
	Dst, _ := base64.URLEncoding.DecodeString(Str)
	if len(Dst) == 0 {
		return 0, string(Dst)
	}
	return len(Dst), string(Dst)
}

func LoadBodyMsg(data []byte, v interface{}) error {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回

	//读取的数据为json格式，需要进行解码
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return  nil
}

func DecodeBodyMsg(v BodyMsg) BodyMsg {
	vr := BodyMsg{}
	vr.Type = v.Type
	_, body := Base64Decode(v.Msg)
	vr.Msg = body
	return vr
}

func LoadBody(data []byte) BodyMsg{
	v := BodyMsg{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	err := LoadBodyMsg(data, &v)
	if err != nil {
		return  v
	}
	v = DecodeBodyMsg(v)
	return  v
}