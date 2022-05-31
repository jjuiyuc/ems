package repository

import (
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

func GetUserByUsername(username string) (user *deremsmodels.User, err error) {
	user, err = deremsmodels.Users(
		qm.Where("username = ?", username),
		qm.Where("deleted_at IS NULL")).One(models.GetDB())
	return
}

func UpdateUser(user *deremsmodels.User) (err error) {
	user.UpdatedAt = null.NewTime(time.Now(), true)
	_, err = user.Update(models.GetDB(), boil.Infer())
	return
}

func InsertLoginLog(loginLog *deremsmodels.LoginLog) (err error) {
	loginLog.CreatedAt = time.Now()
	loginLog.UpdatedAt = null.NewTime(time.Now(), true)
	err = loginLog.Insert(models.GetDB(), boil.Infer())
	return
}
