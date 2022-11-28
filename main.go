package main

import (
	"fmt"
	"ginvue/Server/makeyaml"
)

func main() {
	f, err := makeyaml.ReadCollections("Config/10.0.0.28/mocks/collections.yaml")
	if err != nil {
		fmt.Println(err)
	}

	print(f)

}
