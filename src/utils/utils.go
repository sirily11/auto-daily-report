package utils

import (
	"auto-daily-report/src/config"
	"github.com/spf13/viper"
)

// ReadConfigFromFile reads the configuration from a file.
func ReadConfigFromFile() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.AddConfigPath("/etc/sme-demo/")
	viper.AddConfigPath("$HOME/.sme-demo")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var readConfig config.Config
	err = viper.Unmarshal(&readConfig)
	if err != nil {
		return nil, err
	}

	return &readConfig, nil
}
