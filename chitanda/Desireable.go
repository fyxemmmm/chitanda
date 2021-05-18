package chitanda

import (
	"gopkg.in/yaml.v2"
	"log"
)

type UserConfig map[string]interface{}

type ServerConfig struct {
	Port int32
	Name string
}

type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
}

func InitConfig() *SysConfig  {
	config := &SysConfig{
		Server: &ServerConfig{
			Port: 8080,
		},
	}
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			log.Fatal(err)
		}
	}

	return config
}