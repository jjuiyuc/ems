package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"der-ems/config"
)

var (
	db *sql.DB
)

// Init database
func Init() {
	var err error

	config := config.GetConfig()
	boil.DebugMode = config.GetBool("server.debug")

	db, err = sql.Open(
		config.GetString("db.derems.driver"),
		config.GetString("db.derems.connection"),
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
