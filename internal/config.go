package internal

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".gh-cli")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
