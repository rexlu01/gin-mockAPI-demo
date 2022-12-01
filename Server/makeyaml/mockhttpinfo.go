package makeyaml

import (
	"fmt"
	"ginvue/Server/util"
	"strings"
)

type MockHttpInfo struct {
	SourceIP       string //预留的字段
	IsReturnReal   bool
	UserPath       string
	MockConfigPath string
	RoutesId       string
	VariantId      string
}

//析构函数
func NewMockHttpInfo(SourceIP string) *MockHttpInfo {
	m := &MockHttpInfo{}
	PathNames := make(map[string]string)
	PathNames = GetConfigPath(SourceIP)
	m.UserPath = PathNames["userPath"]
	m.MockConfigPath = PathNames["mockConfigPath"]
	m.IsReturnReal = IsReturnReal(PathNames["mockConfigPath"])
	m.RoutesId, m.VariantId = GetIdsFormCollections(PathNames)

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

	mockConf, err := ReadMocksConfig(mockConfigPath)
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	return *&mockConf.Mock.Returnreal

}

func GetIdsFormCollections(yamlPath map[string]string) (routesId string, variantId string) {

	mockConf, err := ReadMocksConfig(yamlPath["mockConfigPath"])
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	collectionsId := *&mockConf.Mock.Collections.Selected

	collectionsConfList, err := ReadCollections(yamlPath["collectionsPath"])
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}

	for _, collections := range collectionsConfList {
		if collections.Id == collectionsId {
			//这里简单处理 默认列表只有一个元素，后续要处理，为啥会有多个routes？？
			IdStrings := collections.Routes[0]
			routesId = strings.Split(IdStrings, ":")[0]
			variantId = strings.Split(IdStrings, ":")[1]
		}
	}

	return routesId, variantId

}

//获取端口号
func (m *MockHttpInfo) GerPort() int {
	mockConf, err := ReadMocksConfig(m.MockConfigPath)
	if err != nil {
		//预留log模块
		fmt.Println(err)
	}
	return *&mockConf.Server.Port
}

//中间函数，获取Route对象
func (m *MockHttpInfo) GetRoute() (route User) {
	Users, err := ReadUser(m.UserPath)
	if err != nil {
		fmt.Printf("errors : %v", err)
	}

	for _, Route := range Users {
		if Route.Id == m.RoutesId {
			route = Route
		}

	}
	return route

}

//获取HttpURL
func (m *MockHttpInfo) GetUrl() string {
	route := m.GetRoute()
	return route.URL
}

//获取HttpMethod
func (m *MockHttpInfo) GetMethod() string {
	route := m.GetRoute()
	return route.Method
}

//获取HttpRequestHeadres
func (m *MockHttpInfo) GetRequestHeaders() {

}
