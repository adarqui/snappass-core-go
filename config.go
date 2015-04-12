package snappass_core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Debug bool  `json:"debug"`
	Redis Redis `json:"redis"`
	Web   Web   `json:"web"`
}

type Redis struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Db   int    `json:"db"`
	Auth string `json:"auth"`
}

type Web struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Static  string `json:"static"`
	Logging bool   `json:"logging"`
}

func (config *Config) GetBindAddress() string {
	v := fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port)
	return v
}

func ParseConfig(path string) (*Config, error) {
	config := new(Config)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(f, config)
	return config, nil
}
