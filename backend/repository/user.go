package repository

import (
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	deremsmodels "der-ems/models/der-ems"
)

// UserRepository godoc
type UserRepository interface {
	InsertLoginLog(loginLog *deremsmodels.LoginLog) error
	UpdateUser(user *deremsmodels.User) (err error)
	GetLoginLogCount() (int64, error)
	GetUserByUserID(userID int64) (*deremsmodels.User, error)
	GetUserByUsername(username string) (*deremsmodels.User, error)
	GetUserByPasswordToken(token string) (*deremsmodels.User, error)
	GetGroupByGroupID(groupID int64) (*deremsmodels.Group, error)
	GetSubGroupsByGroupID(groupID int64) ([]*deremsmodels.Group, error)
	GetGatewaysPermissionByGroupID(groupID int64) ([]*deremsmodels.GroupGatewayRight, error)
	GetWebpagesPermissionByGroupTypeID(groupTypeID int64) ([]*deremsmodels.GroupTypeWebpageRight, error)
	GetWebpageByWebpageID(webpagesID int64) (*deremsmodels.Webpage, error)
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
	loginLog.UpdatedAt = now
	return loginLog.Insert(repo.db, boil.Infer())
}

// UpdateUser godoc
func (repo defaultUserRepository) UpdateUser(user *deremsmodels.User) (err error) {
	user.UpdatedAt = time.Now().UTC()
	_, err = user.Update(repo.db, boil.Infer())
	return
}

// GetLoginLogCount godoc
func (repo defaultUserRepository) GetLoginLogCount() (int64, error) {
	return deremsmodels.LoginLogs().Count(repo.db)
}

// GetUserByUserID godoc
func (repo defaultUserRepository) GetUserByUserID(userID int64) (*deremsmodels.User, error) {
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

func (repo defaultUserRepository) GetGroupByGroupID(groupID int64) (*deremsmodels.Group, error) {
	return deremsmodels.FindGroup(repo.db, groupID)
}

func (repo defaultUserRepository) GetSubGroupsByGroupID(groupID int64) ([]*deremsmodels.Group, error) {
	return deremsmodels.Groups(
		qm.SQL(`
		WITH RECURSIVE group_path AS
		(
		SELECT *
			FROM der_ems.group
			WHERE id = ?
		UNION ALL
		SELECT g.*
			FROM group_path AS gp JOIN der_ems.group AS g
			ON gp.id = g.parent_id
			AND g.deleted_at IS NULL
		)
		SELECT * FROM group_path;`, groupID)).All(repo.db)
}

func (repo defaultUserRepository) GetGatewaysPermissionByGroupID(groupID int64) ([]*deremsmodels.GroupGatewayRight, error) {
	return deremsmodels.GroupGatewayRights(
		qm.Where("group_id = ?", groupID)).All(repo.db)
}

func (repo defaultUserRepository) GetWebpagesPermissionByGroupTypeID(groupTypeID int64) ([]*deremsmodels.GroupTypeWebpageRight, error) {
	return deremsmodels.GroupTypeWebpageRights(
		qm.Where("type_id = ?", groupTypeID)).All(repo.db)
}

func (repo defaultUserRepository) GetWebpageByWebpageID(webpagesID int64) (*deremsmodels.Webpage, error) {
	return deremsmodels.FindWebpage(repo.db, webpagesID)
}
