package main


import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func ToParam(str string) string {
	return "\"" + str + "\""
}

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1024 * 1024)
	v := LoadBuildInfo("./BuildInfo.json")
	szPathFile := v.PatchDir + v.PatchFile + strconv.Itoa(v.PatchIndex) + v.PatchExt
	szOutputFile := v.OutputDir + v.OutputFile + strconv.Itoa(v.OutputIndex) + v.OutputExt
	_ = os.Remove(szPathFile)
	_ = os.Remove(szOutputFile)
	log.Println("Start Current Client")
	for {
		l, err := client.Read(message)
		log.Println("Get Current Client Message len = ", l, ", err = ", err)
		if err != nil {
			continue
		}
		if l > 0 {
			ve := GetProtocolInfo(message, l)
			log.Println("Get Current Client Type = ", ve.Type)
			if ve.Type == PROTOCOL_INFO_TYPE_START {
				//	文件落地
				err = ioutil.WriteFile(szPathFile, ve.Msg, os.ModeAppend)
				if err != nil {
					log.Println("Get Current Client : Save File Error : ", len(ve.Msg))
					break
				}
			} else {
				break
			}

			_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_START_SUCCESS, []byte("Auto Patch Start")))

			//	d:\SafePatch\Release\AutoPatch.exe -d "d:\SafePatch\patch.txt" -i "E:\WEB\ZooRoot\SogouExplorer" -o "D:\SafePatch\report.3.zip" -p 123456
			cmd := exec.Command(v.AutoPatch, "-d", ToParam(szPathFile), "-i", ToParam(v.InstallDir), "-o", ToParam(szOutputFile), "-p", v.PassWorld )
			log.Println("Patch : ", cmd.Path)
			log.Println("		args : ", cmd.Args)
			err = cmd.Run();
			if err == nil {
				_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_END_SUCCESS, []byte("Auto Patch Success")))
			} else {
				_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_END_FAIL, []byte("Auto Patch Fail : " + err.Error())))
				break
			}

			if err == nil {
				by ,_ := ioutil.ReadFile(szOutputFile)
				_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_REPORT, by))
			}

			_, _ = client.Write(SetProtocolInfo(PROTOCOL_INFO_TYPE_OVER, []byte("Auto Patch Over")))

			break
			//log.Println(message[0:len])
		} else {
			break
		}
	}

	_ = client.Close()
	log.Println("Exit Current Client")
	return nil
}

func main() {

	v := LoadServerInfo("./ServerInfo.json")
	server, err := net.Listen("tcp", v.Ip + ":" + v.Port)
	if err != nil {
		log.Fatal("build start server failed!\n")
		os.Exit(1)
	}
	defer server.Close()

	log.Println("build server is running...")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error\n")
			continue
		}

		log.Println("build client is connectted...")
		go recvMessage(client)
	}
}