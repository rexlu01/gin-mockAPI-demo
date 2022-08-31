package main

import (
	"fmt"
	"ginvue/MockServer/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	yaml := new(util.IPv4Yaml)

	yaml, err := util.ReadYaml("./MockAdmin/Env/", "10.0.0.40")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(yaml.Relus)

	for _, v := range yaml.Relus {
		// var el []map[string]interface{} = make([]map[string]interface{}, 3)
		// v.Env = el
		for _, rspValue := range v.Response {

			for _, v := range rspValue {
				//fmt.Printf("type : %T; value : %v\n", v, v)
				if value, ok := v.(map[interface{}]interface{}); ok {
					for k, v := range value {
						if kv, isE := k.(string); isE {
							fmt.Println(kv)
						}
						if v2, isE := v.(string); isE {
							fmt.Println(v2)
							fmt.Printf("%T", v)
						}

					}

				}
			}
		}
	}

}

func GenerateAPIGET(URL string, RepPer map[string]string, Format string, status int, Body map[string]string) {
	r := gin.Default()
	r.GET(URL, func(c *gin.Context) {
		if strings.Contains(strings.ToLower(Format), "application/json") {
			c.JSON(status, gin.H{})
		}

	})

}
