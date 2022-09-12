package main

import (
	"fmt"
	"ginvue/MockServer/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import (
// 	"ginvue/MockServer/generate"
// )

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password"`
}

func main() {

	//1.获取请求IP

	// 1.创建路由
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	var reqIP string
	// 2.初始化验证
	r.LoadHTMLFiles("./MockAdmin/Web/fromWeb.html")
	r.GET("/UniversalMock", func(c *gin.Context) {
		reqIP = util.GetRequestIP(c)
		fmt.Println(reqIP)
		Conf := make(map[string]string)
		Conf = util.ReadConf("./MockAdmin/Config/mockEnv.conf")
		fmt.Println(Conf)
		IsClint := false
		for k := range Conf {
			if k == reqIP {
				IsClint = true
			}
		}
		if IsClint {
			if isok, err := util.CheckFileName("./MockAdmin/Env/", reqIP); isok && err == nil {
				c.HTML(http.StatusOK, "fromWeb.html", nil)
			} else {
				c.String(http.StatusForbidden, "未找到环境对应的配置文件")
			}

		} else {
			c.String(http.StatusForbidden, "未找到你的环境")
		}

	})
	// JSON绑定
	r.POST("/mockInput", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	//这是提供测试用的真实接口
	// r.GET("/pingTest", func(c *gin.Context) {
	// 	pingAPIdata := map[string]int{
	// 		"success": 0,
	// 		"code":    541994,
	// 	}
	// 	c.JSON(200, pingAPIdata)
	// })

	portNum := util.ReturnPort()
	fmt.Println("portNum", portNum)

	r.Run(reqIP + ":" + fmt.Sprintf("%d", portNum))
}
