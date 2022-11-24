package generate

import (
	"fmt"
	"ginvue/Server/util"
	"strings"

	"github.com/gin-gonic/gin"
)

type Rules struct {
	URL              string
	RequestPerameter []string
	RequestHeader    map[string]string
	Status           int
	ResponseFormat   string
	ResponseBody     map[string]interface{}
}

//总分支处理逻辑
func (r *Rules) TotalMakeYaml(requestIpv4 string) []string {
	var SuccessList []string = []string{}
	//var r Rules
	yaml := new(util.IPv4Yaml)
	yaml, err := util.ReadYaml("./MockAdmin/Env/", requestIpv4)
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
								r.GenerateAPIGET(RulesName)
								SuccessList = append(SuccessList, RulesName)
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
	return SuccessList
}

//组装GET请求
func (r *Rules) GenerateAPIGET(RulesName string) {
	r.RequestHeader = make(map[string]string)
	r.ResponseBody = map[string]interface{}{}
	yaml := new(util.IPv4Yaml)
	yaml, err := util.ReadYaml("./MockAdmin/Env/", "10.0.0.15")
	if err != nil {
		panic(err)
	}

	for _, RelesV := range yaml.Relus {
		if RelesV.Name == RulesName {
			for _, RequeestInfo := range RelesV.Request {
				if URLInterfaceType, ok := RequeestInfo["URL"]; ok {
					if r.URL, ok = URLInterfaceType.(string); !ok {
						break
					} else {
						fmt.Println(r.URL)
					}
				}
				if PerameInterfaceType, ok := RequeestInfo["Perameter"]; ok {
					if Perameter, ok := PerameInterfaceType.(map[interface{}]interface{}); ok {
						for PerameterKey := range Perameter {
							PerameterKeyStr, okKey := PerameterKey.(string)
							if okKey {
								r.RequestPerameter = append(r.RequestPerameter, strings.Split(PerameterKeyStr, ":")...)
								//fmt.Println(RequestPerameter)
							}
						}
					}
				}
				if HeaderInterfaceType, ok := RequeestInfo["Header"]; ok {
					if Header, ok := HeaderInterfaceType.(map[interface{}]interface{}); ok {
						for HeaderKey := range Header {
							HeaderKeyStr, okKey := HeaderKey.(string)
							if okKey {
								HeaderKeySli := strings.Split(HeaderKeyStr, ":")
								r.RequestHeader[HeaderKeySli[0]] = HeaderKeySli[1]
								fmt.Println(r.RequestHeader)
							}

						}

					}
				}
			}
			for _, ResponseInfo := range RelesV.Response {
				if StatusInterface, ok := ResponseInfo["Status"]; ok {
					if r.Status, ok = StatusInterface.(int); !ok {
						break
					}
				}
				if FormmatInterface, ok := ResponseInfo["Format"]; ok {
					if r.ResponseFormat, ok = FormmatInterface.(string); !ok {
						break
					}

				}
				if BodyInterface, ok := ResponseInfo["Body"]; ok {
					for BodyMapKeyInterfaceType, BodyMapValueInterfaceType := range BodyInterface.(map[interface{}]interface{}) {
						BodyMapKeyStr, KeyStrok := BodyMapKeyInterfaceType.(string)
						if KeyStrok {
							r.ResponseBody[BodyMapKeyStr] = BodyMapValueInterfaceType
							fmt.Println(r.ResponseBody)
						} else {
							break
						}

					}

				}

			}

		}

	}

	router := gin.Default()
	router.GET(r.URL, func(c *gin.Context) {
		for _, perameterValue := range r.RequestPerameter {
			c.DefaultQuery(perameterValue, "")
		}
		for HeaderKey, HeaderValue := range r.RequestHeader {
			c.Request.Header.Add(HeaderKey, HeaderValue)
		}

		if strings.Contains(strings.ToLower(r.ResponseFormat), "application/json") {
			c.JSON(r.Status, r.ResponseBody)
		}
	})
	LocalIPv4, err := util.GetLocalIPv4()

	if err != nil {
		//errors.New("本机IP获取失败")
		panic(err)
	}

	port := fmt.Sprintf("%d", util.ReturnPort())

	router.Run(LocalIPv4 + ":" + port)
}

//组装POST请求
