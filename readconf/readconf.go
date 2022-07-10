package readconf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type InterInfoC struct {
	IsMock        bool    `yaml:"IsMock"`
	InterfaceInfo string  `yaml:"InterfaceInfo"`
	Method        string  `yaml:"Method"`
	Fields        []Field `yaml:"Fields"`
	Response      []Resp  `yaml:"Response"`
}

type Field struct {
	Name string              `yaml:"name"`
	Info []map[string]string `yaml:"info"`
}

type Resp struct {
	Name string              `yaml:"name"`
	Info []map[string]string `yaml:"info"`
}

func ReadConf() (in *InterInfoC) {
	in = new(InterInfoC)

	yamlFile, err := ioutil.ReadFile("InterInfoConfig.yaml")

	if err != nil {
		fmt.Printf("%#v", err)
	}
	err = yaml.Unmarshal(yamlFile, in)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	return in

}
