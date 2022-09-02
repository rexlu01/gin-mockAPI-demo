package httpmockclient

import (
	"ginvue/MockServer/generate"
	"ginvue/MockServer/util"
	"time"

	"net/http"
)

func HttpMockAPIclient(requestIpv4 string, rulesName string) {
	var r generate.Rules
	rulesList := r.TotalMakeYaml(requestIpv4)

	yaml := new(util.IPv4Yaml)
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
					for _, rule := range rulesList {
						if rule == RelusV.Name {
							//解析request
							var RequestURL string
							for _, RequestInfoV := range RelusV.Request {
								if URLInterfaceType, ok := RequestInfoV["URL"]; ok {
									if RequestURL, ok = URLInterfaceType.(string); !ok {
										break
									}
								}
								if 

							}

							client := &http.Client{Timeout: time.Duration(10) * time.Second}
							http.NewRequest()

						}
					}

				}
			}
		}
	}
}
