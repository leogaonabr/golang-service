package config

import "github.com/spf13/viper"

// GetPort ...
func GetPort() int {
	return viper.GetInt("PORT")
}

// GetEnv ...
func GetEnv() string {
	return viper.GetString("ENV")
}

// GetVersion ...
func GetVersion() string {
	return "1.0.0"
}
