package main

import (
	"blog/bootstrap"
	"blog/bootstrap/driver"
	"blog/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var app *gin.Engine

// set app run mode
var appEnv map[string]string

func init() {
	appEnv = make(map[string]string)
	appEnv["debug"] = gin.DebugMode
	appEnv["test"] = gin.TestMode
	appEnv["release"] = gin.ReleaseMode
}

func main() {
	fmt.Println("set env run mode")
	config.InitConfig()
	env := config.Config.Section("env")
	mode, err := env.GetKey("Mode")
	if err != nil {
		panic(err)
	}
	// get config port
	port, err := env.GetKey("Port")
	if err != nil {
		panic(err)
	}
	gin.SetMode(appEnv[strings.ToLower(fmt.Sprintf("%s", mode))])

	fmt.Println("set database driver connect")
	// register database connect
	driver.InitConn("mysql")
	// 启用控制台日志颜色
	gin.ForceConsoleColor()

	fmt.Println("set application router")
	// set application router
	app = bootstrap.Init()

	fmt.Printf("The application run at :%s", port)
	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
