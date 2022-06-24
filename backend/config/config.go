package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

// Init config
func Init(dir, env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath(dir)

	if err := config.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "config.ReadInConfig",
			"err":       err,
		}).Fatal()
	}

	if level, err := log.ParseLevel(config.GetString("log.level")); err == nil {
		log.SetLevel(level)
	}
	log.SetOutput(os.Stdout)
}

// GetConfig return config instance
func GetConfig() *viper.Viper {
	return config
}
