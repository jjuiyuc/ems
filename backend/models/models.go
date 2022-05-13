package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
