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
	var config MocksConfig
	file, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

type Collections struct {
	Id     string   `yaml:"id"`
	Routes []string `yaml:"routes"`
}

func ReadCollections(Path string) ([]Collections, error) {
	var coll []Collections
	file, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &coll)
	if err != nil {
		return nil, err
	}

	return coll, err

}

type User struct {
	Id       string `yaml:"id"`
	URL      string `yaml:"url"`
	Method   string `yaml:"method"`
	Variants []Variant
}

type Variant struct {
	ID      string `yaml:"id"`
	Options struct {
		Request struct {
			Headers []string `yaml:"headers"`
			Params  []string `yaml:"params"`
		}
		Response struct {
			Restype string   `yaml:"restype"`
			Status  int      `yaml:"status"`
			Body    []string `yaml:"body"`
		}
	}
}

func ReadUser(path string) ([]User, error) {
	var u []User
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &u)
	if err != nil {
		return nil, err
	}

	return u, err

}
