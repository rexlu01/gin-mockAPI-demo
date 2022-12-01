package util

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"io"
	"io/ioutil"
	"net"
	"os"

	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

const Config string = "Config/"

type IPv4Yaml struct {
	Relus []Rule `yaml:"Rules"`
}

type Rule struct {
	Name     string                   `yaml:"Name"`
	Env      []map[string]interface{} `yaml:"Env"`
	Request  []map[string]interface{} `yaml:"Request"`
	Response []map[string]interface{} `yaml:"Response"`
}

//获取本机IP
func GetLocalIPv4() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		//判断IP是否是回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can not find the client ip address")

}

//获取来源IP
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

//读取Conf
func ReadConf(path string) map[string]string {
	conf := make(map[string]string)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		s := string(b)
		Sli := strings.Split(s, " ")
		conf[Sli[0]] = Sli[1]
	}
	return conf
}

//检查Env下是否包含指定name的文件
func CheckFileName(path string, reqIP string) (bool, error) {
	findFile := path + reqIP + ".yaml"

	_, err := os.Stat(findFile)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err

}

//解析yaml文件
func ReadYaml(path string, reqIP string) (ipYaml *IPv4Yaml, err error) {
	ipYaml = new(IPv4Yaml)
	findFile := path + reqIP + ".yaml"

	yamlFile, err := ioutil.ReadFile(findFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, ipYaml)
	if err != nil {
		return nil, err
	}

	return ipYaml, nil
}

//随机一个端口号
func ReturnPort() (portNo int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2000) + 8000
}

//创建目录及文件
func InitConfigFile(ProbjectId string) {
	//创建文件夹
	err := os.Mkdir("Config/"+ProbjectId, 0777)
	Check(err)
	err = os.Mkdir("Config/"+ProbjectId+"/mocks", 0777)
	Check(err)
	err = os.Mkdir("Config/"+ProbjectId+"/mocks/routes", 0777)
	Check(err)
	// err = os.Mkdir("MocksConfig/"+ProbjectId+"/returntrue", 0777)
	// check(err)

	//创建文件
	f, err := os.Create("Config/" + ProbjectId + "/mocks.config.yaml")
	Check(err)
	defer f.Close()
	f, err = os.Create("Config/" + ProbjectId + "/mocks/collections.yaml")
	Check(err)
	defer f.Close()
	f, err = os.Create("Config/" + ProbjectId + "/mocks/routes/user.yaml")
	Check(err)
	defer f.Close()

	// f, err = os.Create("MocksConfig/" + ProbjectId + "/returntrue/response.yaml")
	// check(err)
	// defer f.Close()

}

//检查错误
func Check(err error) {
	if err != nil {
		panic(err)

	}
}

//判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取某目录下所有的文件夹
func CollectDirs(path string) ([]string, error) {
	var res []string = []string{}
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("read dir fail:%v", err)
		return res, err
	}
	for _, dir := range rd {
		if dir.IsDir() {
			res = append(res, dir.Name())
		}
	}
	return res, nil

}

//验证URL合法
func ValidateURL(u string) (bool, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false, err
	}
	Url, err := url.Parse(u)
	if err != nil || Url.Scheme == "" || Url.Host == "" {
		return false, err
	}

	return true, nil

}
