package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// UserRepository godoc
type UserRepository interface {
	InsertLoginLog(loginLog *deremsmodels.LoginLog) error
	UpdateUser(user *deremsmodels.User) (err error)
	GetLoginLogCount() (int64, error)
	GetUserByUserID(userID int) (*deremsmodels.User, error)
	GetUserByUsername(username string) (*deremsmodels.User, error)
	GetUserByPasswordToken(token string) (*deremsmodels.User, error)
}

type defaultUserRepository struct {
	db *sql.DB
}

// NewUserRepository godoc
func NewUserRepository(db *sql.DB) UserRepository {
	return &defaultUserRepository{db}
}

// InsertLoginLog godoc
func (repo defaultUserRepository) InsertLoginLog(loginLog *deremsmodels.LoginLog) error {
	now := time.Now().UTC()
	loginLog.CreatedAt = now
	loginLog.UpdatedAt = null.NewTime(now, true)
	return loginLog.Insert(repo.db, boil.Infer())
}

// UpdateUser godoc
func (repo defaultUserRepository) UpdateUser(user *deremsmodels.User) (err error) {
	user.UpdatedAt = null.NewTime(time.Now().UTC(), true)
	_, err = user.Update(repo.db, boil.Infer())
	return
}

// GetLoginLogCount godoc
func (repo defaultUserRepository) GetLoginLogCount() (int64, error) {
	return deremsmodels.LoginLogs().Count(repo.db)
}

// GetUserByUserID godoc
func (repo defaultUserRepository) GetUserByUserID(userID int) (*deremsmodels.User, error) {
	return deremsmodels.FindUser(repo.db, userID)
}

// GetUserByUsername godoc
func (repo defaultUserRepository) GetUserByUsername(username string) (*deremsmodels.User, error) {
	return deremsmodels.Users(
		qm.Where("username = ?", username),
		qm.Where("deleted_at IS NULL")).One(repo.db)
}

// GetUserByPasswordToken godoc
func (repo defaultUserRepository) GetUserByPasswordToken(token string) (*deremsmodels.User, error) {
	return deremsmodels.Users(
		qm.Where("reset_pwd_token = ?", token),
		qm.Where("pwd_token_expiry > ?", time.Now().UTC())).One(repo.db)
}
