package internal

import "github.com/spf13/viper"

func ReadVal() (string, error) {
	viperConfig := *viper.New()
	viperConfig.SetConfigFile("config.yaml")
	err := viperConfig.ReadInConfig() // Find and read the config file
	if err != nil {
		return "", err
	}
	return viperConfig.GetString("some_key"), nil
}
