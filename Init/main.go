package main

import (
	"fmt"
	"ginvue/Server/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	LocalIP, err := util.GetLocalIPv4()
	if err != nil {
		panic("获取本地IP失败")
	}

	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("/initializeMock", func(ctx *gin.Context) {
		SourceIp := util.GetRequestIP(ctx)
		//获取当前目录
		//检查以SourceIP命名的目录是否存在
		pwd, _ := os.Getwd()
		fmt.Printf("pwd:%s", pwd)
		path := util.Config + SourceIp
		fmt.Printf("path:%s\n", path)
		if ok, _ := util.PathExists(path); ok {
			ctx.String(http.StatusOK, "工程目录已经初始化过，不必重复请求")
		} else {
			util.InitConfigFile(SourceIp)
			ctx.String(http.StatusOK, "工程目录初始化成功")
		}
	})

	portNum := "9615"
	//fmt.Println("portNum", portNum)

	r.Run(LocalIP + ":" + portNum)
}
