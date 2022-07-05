package models

import (
	"database/sql"

	// import the driver
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db *sql.DB
)

// Init init database
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

// Close godoc
func Close() {
	db.Close()
}

// GetDB godoc
func GetDB() *sql.DB {
	return db
}
