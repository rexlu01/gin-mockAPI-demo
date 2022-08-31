package main

import (
	"fmt"
	"ginvue/MockServer/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	yaml := new(util.IPv4Yaml)
	yaml, err := util.ReadYaml("./MockAdmin/Env/", "10.0.0.15")
	if err != nil {
		panic(err)
	}
	var RulesName string
	for _, RelusV := range yaml.Relus {
		for _, EnvV := range RelusV.Env {
			//fmt.Println(EnvV)
			//EnvV = make(map[string]interface{})
			if ReturnRealEnv, ok := EnvV["ReturnRealEnv"]; ok {
				//fmt.Println(ReturnRealEnv)
				if ReturnRealEnv.(bool) {
					//如果是真的环境todo
				} else {
					//如果是mock环境
					//拿到规则名字
					RulesName = RelusV.Name
					for _, ReqV := range RelusV.Request {
						if ReqMethodV, ok := ReqV["Method"]; ok {
							if strings.ToLower(ReqMethodV.(string)) == "get" {
								fmt.Println(RulesName)
								//todo这里根据规则name生成对应的Get mock API
							} else if strings.ToLower(ReqMethodV.(string)) == "post" {
								fmt.Println(RulesName)
								//todo这里根据规则name生成对应的Post mock API
							}

						}

					}
				}

			}
		}

	}

	fmt.Println(util.ReturnPort())

}

func GenerateAPIGET(Name string, URL string, ReqHeader map[string]string, ReqPer map[string]string, Format string, status int, Body map[string]string) {
	r := gin.Default()
	r.GET(URL, func(c *gin.Context) {
		if strings.Contains(strings.ToLower(Format), "application/json") {
			ReqHeader = make(map[string]string)
			for k, v := range ReqHeader {
				c.Header(k, v)
			}
			ReqPer = make(map[string]string)
			for k, v := range ReqPer {
				c.DefaultQuery(k, v)
			}
			c.JSON(status, Body)
		}
	})

}
