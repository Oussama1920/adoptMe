package config

import (
	"github.com/spf13/viper"
)

// Init setup the basic environment for the service
func Init(configName string, configPaths []string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	for _, configPath := range append(configPaths, ".") {
		viper.AddConfigPath(configPath)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// GetDataConfiguration unmarshal struct data from configuration
func GetDataConfiguration(key string, serviceConf interface{}) error {
	return viper.UnmarshalKey(key, serviceConf)
}
