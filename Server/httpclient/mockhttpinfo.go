package httpclient

import (
	"fmt"
	"ginvue/Server/makeyaml"
	"ginvue/Server/util"
)

type MockHttpInfo struct {
	IsReturnReal bool
	UserPath     string
	VariantId    string
}

//析构函数
func NewMockHttpInfo(SourceIP string) *MockHttpInfo {
	m := &MockHttpInfo{}
	PathNames := make(map[string]string)
	PathNames = GetConfigPath(SourceIP)
	m.UserPath = PathNames["userPath"]
	m.IsReturnReal = IsReturnReal(PathNames["mockConfigPath"])

	return m

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

func IsReturnReal(mockConfigPath string) bool {

	mockConf, err := makeyaml.ReadMocksConfig(mockConfigPath)
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	return *&mockConf.Mock.Returnreal

}

func GetVariantId(yamlPath map[string]string) (routesId string, variantId string) {

	mockConf, err := makeyaml.ReadMocksConfig(yamlPath["mockConfigPath"])
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	collectionsId := *&mockConf.Mock.Collections.Selected

	collectionsConfList, err := makeyaml.ReadCollections(yamlPath["collectionsPath"])
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	for _, collections := range collectionsConfList {
		if collections.Id == collectionsId {

		}
	}

}
