package main

import (
	"fmt"
	"ginvue/MockServer/util"
	"net/http"

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
		SourceIP := util.GetRequestIP(ctx)
		fmt.Println(SourceIP)
		//获取当前目录
		//检查以SourceIP命名的目录是否存在
		path := "MocksConfig/" + SourceIP
		if ok, _ := util.PathExists(path); ok {
			ctx.String(http.StatusOK, "工程目录已经初始化过，不必重复请求")
		} else {
			util.InitConfigFile(SourceIP)
			ctx.String(http.StatusOK, "工程目录初始化成功")
		}
	})

	portNum := util.ReturnPort()
	//fmt.Println("portNum", portNum)

	r.Run(LocalIP + ":" + fmt.Sprintf("%d", portNum))
}
