package main

import (
	"mockapi/readconf"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenGinPost(reqInfo *readconf.InterInfoC) {

	// ClientIp, err := genutil.GetClientIp()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if reqInfo.IsMock {
		apiInfo := reqInfo.InterfaceInfo
		//获取路由地址
		routeURI := strings.Split(apiInfo, " ")[1]
		//获取请求参数并且初始化
		param := make(map[string]interface{})
		for _, pInfoS := range reqInfo.Fields {
			for _, k := range pInfoS.Info {
				param[pInfoS.Name] = k["rdefault"]
			}
		}

		//获取响应参数并且初始化（后期动态参数）
		ResponseP := make(map[string]interface{})
		for _, ResPlist := range reqInfo.Response {
			for _, k := range ResPlist.Info {
				ResponseP[ResPlist.Name] = k["default"]
			}

		}

		r := gin.Default()

		r.POST(routeURI, func(c *gin.Context) {

			if err := c.ShouldBindJSON(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, ResponseP)
		})

		r.Run(":7090")

	}

}

// func GenGinGet() {
// 	r := gin.Default()
// }

func main() {
	YamlContent := readconf.ReadConf()

	//fmt.Println(YamlContent.Fields)

	GenGinPost(YamlContent)

}
