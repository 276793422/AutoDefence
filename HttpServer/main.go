package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	//v := LoadConfig("./config.json")
	//fmt.Println(v)
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/PatchMsg.zoo", Param)

	// Start server
	str := "0.0.0.0" + ":" + "20202"
	log.Println("ip:port = ", str)
	e.Logger.Fatal(e.Start(str))
}

// Handler
func Param(c echo.Context) error {
	req := c.Request()

	body ,err := ioutil.ReadAll(req.Body)
	if err != nil {
		return  err
	}

	v := LoadBody(body)

	_ = c.String(http.StatusOK, "Success")
	if v.Type != "Zoo" || v.Msg == "" {
		log.Println("From :" + c.RealIP() + " : error ")
		log.Println("v.Type = [" + v.Type + "]" )
		log.Println("v.Msg len = " + strconv.Itoa(len(v.Msg)))
		return nil		//	出错了
	}
	vp := LoadConfig("./HttpConfig.json")

	_ = os.Remove(vp.PatchFileName)												//	删了文件
	_ = ioutil.WriteFile(vp.PatchFileName, []byte(v.Msg), os.ModeAppend)		//	补上文件

	log.Println("Start Call Client ...")
	cmd := exec.Command(vp.ClientPath, vp.PatchFileName)
	err = cmd.Run();
	if err != nil {
		log.Println("Call Client Over ... error = ", err.Error())
	} else {
		log.Println("Call Client Over ... no error ")
	}

	return err
}