package config

import (
	log "github.com/sirupsen/logrus"
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
		log.Fatal("err ReadInConfig: ", err)
	}

	if level, err := log.ParseLevel(config.GetString("log.level")); err == nil {
		log.SetLevel(level)
	}
}

// GetConfig return config instance
func GetConfig() *viper.Viper {
	return config
}
