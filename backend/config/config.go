package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	config    *viper.Viper
	gitBranch string
	gitCommit string
)

// Init config
func Init(dir, env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath(dir)

	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

// GetConfig return config instance
func GetConfig() *viper.Viper {
	return config
}
