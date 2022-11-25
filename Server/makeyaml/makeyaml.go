package makeyaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MocksConfig struct {
	Log    string `yaml:"log"`
	Server struct {
		Port int `yaml:"port"`
	}
	Mock struct {
		Returnreal  bool `yaml:"returnreal"`
		Collections struct {
			Selected string `yaml:"selected"`
		}
	}
}

func ReadMocksConfig(Path string) (*MocksConfig, error) {
	var config *MocksConfig
	file, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, err
}
