package config

import (
	"bytes"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config    *viper.Viper
	gitBranch string
	gitCommit string
)

// Notes: https://github.com/sirupsen/logrus/issues/444
type defaultLogger struct {
	formatter log.Formatter
}

// OutputSplitter godoc
type OutputSplitter struct{}

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
	log.SetFormatter(defaultLogger{formatter: log.StandardLogger().Formatter})
	log.SetOutput(&OutputSplitter{})
}

// GetConfig return config instance
func GetConfig() *viper.Viper {
	return config
}

func (l defaultLogger) Format(entry *log.Entry) ([]byte, error) {
	entry.Data["branch"] = gitBranch
	entry.Data["commit"] = gitCommit
	return l.formatter.Format(entry)
}

func (splitter *OutputSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}
