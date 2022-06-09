package repository

import (
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

// UserRepository ...
type UserRepository interface {
	GetUserByUsername(username string) (user *deremsmodels.User, err error)
	UpdateUser(user *deremsmodels.User) (err error)
	InsertLoginLog(loginLog *deremsmodels.LoginLog) (err error)
	GetProfileByUserID(userID int) (user *deremsmodels.User, err error)
}

type defaultUserRepository struct{}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return &defaultUserRepository{}
}

// GetUserByUsername ...
func (repo defaultUserRepository) GetUserByUsername(username string) (*deremsmodels.User, error) {
	return deremsmodels.Users(
		qm.Where("username = ?", username),
		qm.Where("deleted_at IS NULL")).One(models.GetDB())
}

// UpdateUser ...
func (repo defaultUserRepository) UpdateUser(user *deremsmodels.User) (err error) {
	user.UpdatedAt = null.NewTime(time.Now(), true)
	_, err = user.Update(models.GetDB(), boil.Infer())
	return
}

// InsertLoginLog ...
func (repo defaultUserRepository) InsertLoginLog(loginLog *deremsmodels.LoginLog) error {
	loginLog.CreatedAt = time.Now()
	loginLog.UpdatedAt = null.NewTime(time.Now(), true)
	return loginLog.Insert(models.GetDB(), boil.Infer())
}

// GetProfileByUserID ...
func (repo defaultUserRepository) GetProfileByUserID(userID int) (*deremsmodels.User, error) {
	return deremsmodels.FindUser(models.GetDB(), userID)
}
