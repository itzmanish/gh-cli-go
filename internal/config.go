package internal

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".gh-cli")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func SetConfig(data map[string]interface{}, filename string, path string) error {
	for k, v := range data {
		viper.Set(k, v)
	}
	if path != "" {
		return viper.WriteConfigAs(path + "/" + filename)
	}
	config_dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	return viper.WriteConfigAs(config_dir + "/" + filename)

}
