package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type BuildInfo struct {
	//	功能exe
	AutoPatch	string			//	默认exe

	//	Data 数据路径	-d 参数
	PatchDir	string			//	patch 文件目录
	PatchFile	string			//	patch 文件名字
	PatchIndex	int				//	patch 文件索引
	PatchExt	string			//	patch 文件后缀

	//	InstallDir 安装包文件路径	-i 参数
	InstallDir	string			//	安装包路径

	//	导出结果目录	-o
	OutputDir	string			//	Output 文件目录
	OutputFile	string			//	Output 文件名字
	OutputIndex	int				//	Output 文件索引
	OutputExt	string			//	Output 文件后缀

	//	默认密码	-p
	PassWorld	string			//	密码
}

func _LoadBuildInfo(filename string, v interface{}) {
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

func SaveBuildInfo(filename string, v interface{}) bool {
	li, _ := json.Marshal(v)
	err := ioutil.WriteFile(filename, li, os.ModeAppend)
	if err != nil {
		return false
	}
	return  true
}

//	"./BuildInfo.json"
func LoadBuildInfo(filename string) BuildInfo{
	v := BuildInfo{}
	//	取HID，如果取不到就生成一个，
	_LoadBuildInfo(filename, &v)
	if v.AutoPatch == "" {
		v.AutoPatch = "D:\\SafePatch\\Release\\AutoPatch.exe"
	}
	if v.PatchDir == "" && v.PatchFile == "" && v.PatchIndex == 0 && v.PatchExt == ""{
		v.PatchDir = GetCurrentPath() + "Patch\\"
		v.PatchFile = "Patch."
		v.PatchIndex = 0
		v.PatchExt = ".txt"
	}
	if v.InstallDir == ""{
		v.InstallDir = "E:\\WEB\\ZooRoot\\SogouExplorer"
	}
	if v.OutputDir == "" && v.OutputFile == "" && v.OutputIndex == 0 && v.OutputExt == ""{
		v.OutputDir = GetCurrentPath() + "Repair\\"
		v.OutputFile = "Repair."
		v.OutputIndex = 0
		v.OutputExt = ".zip"
	}
	if v.PassWorld == ""{
		v.PassWorld = "123456"
	}

	v.PatchIndex ++
	v.OutputIndex ++

	SaveBuildInfo(filename, v)
	return v
}
