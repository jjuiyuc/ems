package testutils

import (
	"path/filepath"
	"runtime"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/testutils/fixtures"
)

func GetConfigDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "config")
}

func SeedUtUser() {
	models.GetDB().Exec("truncate table user")
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(fixtures.UtUser.Password), bcrypt.DefaultCost)
	user := &deremsmodels.User{
		Username:       fixtures.UtUser.Username,
		Password:       string(hashPassword[:]),
		ExpirationDate: fixtures.UtUser.ExpirationDate,
	}
	user.Insert(models.GetDB(), boil.Infer())
}
