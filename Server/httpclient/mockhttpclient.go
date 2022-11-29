package httpclient

import (
	"fmt"
	"ginvue/Server/util"
)

type req struct {
	Method  string
	UrI     string
	Headers string
	Params  []byte
}

func GetConfigPath(SourceIP string) (yamlPath map[string]string) {
	yamlPath = make(map[string]string)
	yamlPath["mockConfigPath"] = ""
	yamlPath["collectionsPath"] = ""
	yamlPath["userPath"] = ""

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	DirList, _ := util.CollectDirs("Config/")

	//这里是和来源IP做匹配
	for _, collSourceIP := range DirList {
		if collSourceIP == SourceIP {
			yamlPath["mockConfigPath"] = util.Config + collSourceIP + "/mocks.config.yaml"
			yamlPath["collectionsPath"] = util.Config + collSourceIP + "/mocks/collections.yaml"
			yamlPath["userPath"] = util.Config + collSourceIP + "/mocks/routes/user.yaml"
		}

	}

	return yamlPath

}

func IsReturnMock() bool {

}
