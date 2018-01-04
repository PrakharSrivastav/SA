package config

import (
	"github.com/spf13/viper"
)

// Config reads the configurations from the json file and return the map
func Config() (map[string]interface{}, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viper.AllSettings(), nil
}
