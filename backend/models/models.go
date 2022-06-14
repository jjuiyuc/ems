package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db *sql.DB
)

// Init database
func Init(cfg *viper.Viper) {
	var err error

	boil.DebugMode = cfg.GetBool("server.debug")

	db, err = sql.Open(
		cfg.GetString("db.derems.driver"),
		cfg.GetString("db.derems.connection"),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "sql.Open",
			"err":       err,
		}).Panic()
	}
}

func Close() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}
