package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in (root DIR)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalf("Error reading config file: %+v", err)
	}

	viper.WatchConfig() // Watch the config file for changes
}

// GetString is a wrapper for viper
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt is a wrapper for viper
func GetInt(key string) int {
	return viper.GetInt(key)
}
