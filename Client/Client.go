package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var nExit int = 0
	v := LoadServerInfo("./ServerConfig.json")
	client, err := net.Dial("tcp", v.Ip + ":" + v.Port)
	log.Println("Client is Running... : ", "\"" + v.Ip + ":" + v.Port + "\"")
	if err != nil {
		log.Println("Client is dailing failed!")
		os.Exit(1)
	}

	log.Println("Client is Running...")

	var strPatch string = "./patch.txt"
	if len(os.Args) > 1 && os.Args[1] != "" {
		strPatch = os.Args[1]
	}
	data, err := ioutil.ReadFile(strPatch)

	log.Println("Send Patch File : ", len(data))
	_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_START, data))
	var msg []byte
	msg = make([]byte, 10*1024*1024)
	for ; ; {
		l ,_ := client.Read(msg)
		if l > 0 {
			vr := GetProtocolInfo(msg, l)
			switch vr.Type {
			case PROTOCOL_INFO_TYPE_START_SUCCESS:
				log.Println("request : Start Success : ", string(vr.Msg))
				break
			case PROTOCOL_INFO_TYPE_MSG:
				log.Println("request : Msg : ", string(vr.Msg))
				break
			case PROTOCOL_INFO_TYPE_END_SUCCESS:
				log.Println("request : End Success : ", string(vr.Msg))
				break
			case PROTOCOL_INFO_TYPE_END_FAIL:
				str := strings.Split(string(vr.Msg), " ")
				for i := 0; i < len(str) ; i++ {
					nRet, err := strconv.Atoi(str[i])
					if err == nil {
						nExit = nRet
					}
				}
				log.Println("request : End Fail : ", string(vr.Msg))
				break
			case PROTOCOL_INFO_TYPE_REPORT:
				log.Println("request : Report : Save File To Local")
				if v.Report == "" {
					v.Report = "patch.report.zip"
				}
				_ = ioutil.WriteFile(v.Report, vr.Msg, os.ModeAppend)
				break
			default:
				break
			}
			if vr.Type == PROTOCOL_INFO_TYPE_OVER {
				log.Println("request : Report : ", string(vr.Msg))
				_ = client.Close()
				break
			}
		} else {
			break
		}
	}

	ReportEmail(nExit, v)
	log.Println("Client will exit...")
	os.Exit(nExit)
}

func ReportEmail(nExit int, v ServerInfo){
	if nExit == 0{
		cmd := exec.Command(v.Email, "0", "获取数据成功", v.Report)
		_ = cmd.Run()
	} else {
		cmd := exec.Command(v.Email, strconv.Itoa(nExit), "获取数据失败")
		_ = cmd.Run()
	}
}

