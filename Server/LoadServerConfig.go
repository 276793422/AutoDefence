package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ServerInfo struct {
	Ip		string
	Port	string
}

func _LoadServerInfo(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func SaveServerInfo(filename string, v interface{}) bool {
	li, _ := json.Marshal(v)
	err := ioutil.WriteFile(filename, li, os.ModeAppend)
	if err != nil {
		return false
	}
	return  true
}

//	"./ServerInfo.json"
func LoadServerInfo(filename string) ServerInfo{
	v := ServerInfo{}
	//	取HID，如果取不到就生成一个，
	_LoadServerInfo(filename, &v)
	if v.Ip == "" && v.Port == "" {
		v.Ip = "0.0.0.0"
		v.Port = "20200"

		SaveServerInfo(filename, v)
	}
	return v
}
