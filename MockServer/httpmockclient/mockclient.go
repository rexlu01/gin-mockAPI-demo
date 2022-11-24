package httpmockclient

import (
	"encoding/json"
	"fmt"
	"ginvue/MockServer/generate"
	"ginvue/MockServer/util"
	"io/ioutil"
	"strings"
	"time"

	"net/http"
)

func HttpMockAPIclient(requestIpv4 string, rulesName string) string {
	var r generate.Rules
	rulesList := r.TotalMakeYaml(requestIpv4)

	yaml := new(util.IPv4Yaml)
	for _, RelusV := range yaml.Relus {
		for _, EnvV := range RelusV.Env {
			//fmt.Println(EnvV)
			//EnvV = make(map[string]interface{})
			if ReturnRealEnv, ok := EnvV["ReturnRealEnv"]; ok {
				//fmt.Println(ReturnRealEnv)
				if ReturnRealEnv.(bool) == false {
					//如果是mock环境，就生成一个http API（用Gin）
				} else {
					//如果是真的环境todo,这里应该是转发请求
					//拿到规则名字
					for _, rule := range rulesList {
						if rule == RelusV.Name {
							//解析request
							var RequestURL string
							var Method string
							var HeaderKeySli []string = []string{}
							var PerameterKeyScl []string = []string{}
							for _, RequestInfoV := range RelusV.Request {
								if URLInterfaceType, ok := RequestInfoV["URL"]; ok {
									if RequestURL, ok = URLInterfaceType.(string); !ok {
										break
									}
								}
								if MethodInterfaceType, ok := RequestInfoV["Method"]; ok {
									if Method, ok = MethodInterfaceType.(string); !ok {
										break
									}
								}

								if HeaderInterfaceType, ok := RequestInfoV["Header"]; ok {
									if Header, ok := HeaderInterfaceType.(map[interface{}]interface{}); ok {
										for HeaderKey := range Header {
											HeaderKeyStr, okKey := HeaderKey.(string)
											if okKey {
												HeaderKeySli = strings.Split(HeaderKeyStr, ":")
											}

										}

									}
								}

								if PerameInterfaceType, ok := RequestInfoV["Perameter"]; ok {
									if Perameter, ok := PerameInterfaceType.(map[interface{}]interface{}); ok {
										for PerameterKey := range Perameter {
											PerameterKeyStr, okKey := PerameterKey.(string)
											if okKey {
												PerameterKeyScl = strings.Split(PerameterKeyStr, ":")
											}
										}
									}
								}

								client := &http.Client{Timeout: time.Duration(10) * time.Second}

								reqset, err := http.NewRequest(strings.ToUpper(Method), RequestURL, nil)
								reqset.Header.Add(HeaderKeySli[0], HeaderKeySli[1])
								reqset.URL.Query().Add(PerameterKeyScl[0], PerameterKeyScl[1])

								if err != nil {
									panic(err)
								}

								resp, err := client.Do(reqset)
								if err != nil {
									fmt.Printf("%v", err)
									return ""
								}

								defer resp.Body.Close()
								body, err := ioutil.ReadAll(resp.Body)
								if err != nil {
									return ""
								}
								//这里还要解析json
								BodyRes, err := json.Marshal(body)
								if err != nil {
									return string(BodyRes)
								}
								return ""

							}
						}

					}
				}
			}
		}
	}

	return ""
}
