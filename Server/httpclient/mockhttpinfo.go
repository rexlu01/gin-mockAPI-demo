package httpclient

import (
	"fmt"
	"ginvue/Server/makeyaml"
	"ginvue/Server/util"
)

type MockHttpInfo struct {
	IsReturnReal    bool
	MockConfigPath  string
	CollectionsPath string
	UserPath        string
	CollectionId    string
	RouteId         string
	variantId       string
}

//析构函数
func NewMockHttpInfo(SourceIP string) *MockHttpInfo {
	m := &MockHttpInfo{}
	PathNames := make(map[string]string)
	PathNames = GetConfigPath(SourceIP)
	m.MockConfigPath = PathNames["mockConfigPath"]
	m.CollectionsPath = PathNames["collectionsPath"]
	m.UserPath = PathNames["userPath"]

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

func IsReturnReal(SourceIP string) bool {
	//yamlPath := GetConfigPath(SourceIP)

	mockConf, err := makeyaml.ReadMocksConfig(yamlPath["mockConfigPath"])
	if err != nil {
		fmt.Println(err)
	}

	return *&mockConf.Mock.Returnreal

}
