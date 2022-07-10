package genutil

import (
	"errors"
	"net"
)

//获取本机Ip
func GetClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		//检查IP地址判断是否为回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}

	}
	return "", errors.New("Can not find the client ip address!")

}

// func main() {
// 	clientIp, err := GetClientIp()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("%s\n", clientIp)
// }
