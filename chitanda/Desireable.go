package chitanda

import (
	"gopkg.in/yaml.v2"
	"log"
)

//递归读取用户配置文件
func GetConfigValue(m UserConfig,prefix []string,index int) interface{}  {
	key:=prefix[index]
	if v,ok:=m[key];ok{
		if index==len(prefix)-1{ //到了最后一个
			return v
		}else{
			index=index+1
			if mv,ok:=v.(UserConfig);ok{ //值必须是UserConfig类型
				return GetConfigValue(mv,prefix,index)
			}else{
				return  nil
			}

		}
	}
	return  nil
}

type UserConfig map[interface{}]interface{}

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