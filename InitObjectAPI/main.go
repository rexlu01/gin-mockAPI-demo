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
		ctx.String(http.StatusOK, "initialize finished")
	})

	portNum := util.ReturnPort()
	//fmt.Println("portNum", portNum)

	r.Run(LocalIP + ":" + fmt.Sprintf("%d", portNum))
}
