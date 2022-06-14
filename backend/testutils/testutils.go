package testutils

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"

	deremsmodels "der-ems/models/der-ems"
	"der-ems/testutils/fixtures"
)

func GetConfigDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "config")
}

func SeedUtUser(db *sql.DB) (err error) {
	_, err = db.Exec("truncate table user")
	if err != nil {
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(fixtures.UtUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := &deremsmodels.User{
		Username:       fixtures.UtUser.Username,
		Password:       string(hashPassword[:]),
		ExpirationDate: fixtures.UtUser.ExpirationDate,
	}
	err = user.Insert(db, boil.Infer())
	return
}

func GetAuthorization(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
