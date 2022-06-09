package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// UserRepository ...
type UserRepository interface {
	InsertLoginLog(loginLog *deremsmodels.LoginLog) error
	UpdateUser(user *deremsmodels.User) (err error)
	GetLoginLogCount() (int64, error)
	GetProfileByUserID(userID int) (*deremsmodels.User, error)
	GetUserByUsername(username string) (*deremsmodels.User, error)
}

type defaultUserRepository struct {
	db *sql.DB
}

// NewUserRepository ...
func NewUserRepository(db *sql.DB) UserRepository {
	return &defaultUserRepository{db}
}

// InsertLoginLog ...
func (repo defaultUserRepository) InsertLoginLog(loginLog *deremsmodels.LoginLog) error {
	loginLog.CreatedAt = time.Now()
	loginLog.UpdatedAt = null.NewTime(time.Now(), true)
	return loginLog.Insert(repo.db, boil.Infer())
}

// UpdateUser ...
func (repo defaultUserRepository) UpdateUser(user *deremsmodels.User) (err error) {
	user.UpdatedAt = null.NewTime(time.Now(), true)
	_, err = user.Update(repo.db, boil.Infer())
	return
}

// GetLoginLogCount ...
func (repo defaultUserRepository) GetLoginLogCount() (int64, error) {
	return deremsmodels.LoginLogs().Count(repo.db)
}

// GetProfileByUserID ...
func (repo defaultUserRepository) GetProfileByUserID(userID int) (*deremsmodels.User, error) {
	return deremsmodels.FindUser(repo.db, userID)
}

// GetUserByUsername ...
func (repo defaultUserRepository) GetUserByUsername(username string) (*deremsmodels.User, error) {
	return deremsmodels.Users(
		qm.Where("username = ?", username),
		qm.Where("deleted_at IS NULL")).One(repo.db)
}
