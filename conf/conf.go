package conf

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/vsixz/prego/log"
)

type config struct {
	viper.Viper
}

type ConfigType string

const (
	JsonConfigType ConfigType = "json"
	TomlConfigType ConfigType = "toml"
	YamlConfigType ConfigType = "yaml"
	YmlConfigType  ConfigType = "yml"
	IniConfigType  ConfigType = "ini"
	EnvConfigType  ConfigType = "env"
)

func Read(file string, configType ...ConfigType) *config {
	if len(configType) > 0 {
		viper.SetConfigType(string(configType[0]))
	}
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("read config failed: %s", err)
	}
	v := viper.GetViper()
	return &config{
		Viper: *v,
	}
}

func (c config) GetOrDefault(key string, def interface{}) interface{} {
	if !c.IsSet(key) {
		return def
	}
	return c.Get(key)
}

func (c config) GetUint8(key string) uint8 {
	return cast.ToUint8(c.Get(key))
}
