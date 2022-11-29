package main

import (
	"fmt"
	"ginvue/Server/makeyaml"
)

func main() {
	f, err := makeyaml.ReadUser("Config/10.0.0.28/mocks/routes/user.yaml")
	if err != nil {
		fmt.Println(err)
	}

	for i := range *f {
		fmt.Println((*f)[i])
	}

}
