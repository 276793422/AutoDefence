package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)


//定义配置文件解析后的结构
type Config struct {
	ClientPath		string		//	客户端路径
	PatchFileName	string		//	Patch 文件名字
}

func LoadConfigInfo(filename string, v interface{}) {
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

func SaveConfigInfo(filename string, v interface{}) bool {
	li, _ := json.Marshal(v)
	err := ioutil.WriteFile(filename, li, os.ModeAppend)
	if err != nil {
		return false
	}
	return  true
}

func LoadConfig(filename string) Config{
	v := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	LoadConfigInfo(filename, &v)
	if v.ClientPath == "" || v.PatchFileName == "" {
		if v.ClientPath == "" {
			v.ClientPath = "go_build_Client_.exe"
		}
		if v.PatchFileName == "" {
			v.PatchFileName = "patch.txt"
		}
		SaveConfigInfo(filename, v)
	}
	return  v
}