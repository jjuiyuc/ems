package repository

import (
	"database/sql"
	"fmt"
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
	GetUserByUserID(userID int64) (*deremsmodels.User, error)
	GetUserByUsername(username string) (*deremsmodels.User, error)
	GetUserByPasswordToken(token string) (*deremsmodels.User, error)
	GetUserCountByGroupID(groupID int64) (int64, error)
	CreateGroup(group *deremsmodels.Group) (err error)
	UpdateGroup(group *deremsmodels.Group) (err error)
	DeleteGroup(userID, groupID int64) (err error)
	IsGroupNameExistedOnSameLevel(group *deremsmodels.Group) bool
	GetGroupByGroupID(groupID int64) (*deremsmodels.Group, error)
	GetGroupsByGroupID(groupID int64) ([]*deremsmodels.Group, error)
	GetGroupTypes() ([]*deremsmodels.GroupType, error)
	GetGatewaysPermissionByGroupID(groupID int64, findDisabled bool) ([]*deremsmodels.GroupGatewayRight, error)
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

func (repo defaultUserRepository) GetUserCountByGroupID(groupID int64) (int64, error) {
	return deremsmodels.Users(
		qm.Where("group_id = ?", groupID),
		qm.Where("deleted_at IS NULL")).Count(repo.db)
}

func (repo defaultUserRepository) CreateGroup(group *deremsmodels.Group) (err error) {
	return group.Insert(repo.db, boil.Infer())
}

func (repo defaultUserRepository) UpdateGroup(group *deremsmodels.Group) (err error) {
	group.UpdatedAt = time.Now().UTC()
	_, err = group.Update(repo.db, boil.Infer())
	return
}

func (repo defaultUserRepository) DeleteGroup(userID, groupID int64) (err error) {
	group, err := deremsmodels.FindGroup(repo.db, groupID)
	if err == nil {
		now := time.Now().UTC()
		group.UpdatedAt = now
		group.DeletedAt = null.TimeFrom(now)
		group.DeletedBy = null.Int64From(userID)
		_, err = group.Update(repo.db, boil.Infer())
	}
	return
}

func (repo defaultUserRepository) IsGroupNameExistedOnSameLevel(group *deremsmodels.Group) (exist bool) {
	exist, _ = deremsmodels.Groups(
		qm.Where("name = ?", group.Name),
		qm.Where("parent_id = ?", group.ParentID),
		qm.Where("deleted_at IS NULL")).Exists(repo.db)
	return
}

func (repo defaultUserRepository) GetGroupByGroupID(groupID int64) (*deremsmodels.Group, error) {
	return deremsmodels.FindGroup(repo.db, groupID)
}

func (repo defaultUserRepository) GetGroupsByGroupID(groupID int64) ([]*deremsmodels.Group, error) {
	return deremsmodels.Groups(
		qm.SQL(fmt.Sprintf(`
		WITH RECURSIVE group_path AS
		(
		SELECT *
			FROM %s
			WHERE id = ?
		UNION ALL
		SELECT g.*
			FROM group_path AS gp JOIN %s AS g
			ON gp.id = g.parent_id
			AND g.deleted_at IS NULL
		)
		SELECT * FROM group_path;`, "`group`", "`group`"), groupID)).All(repo.db)
}

func (repo defaultUserRepository) GetGroupTypes() ([]*deremsmodels.GroupType, error) {
	return deremsmodels.GroupTypes().All(repo.db)
}

func (repo defaultUserRepository) GetGatewaysPermissionByGroupID(groupID int64, findDisabled bool) ([]*deremsmodels.GroupGatewayRight, error) {
	if findDisabled {
		return deremsmodels.GroupGatewayRights(
			qm.Where("group_id = ?", groupID)).All(repo.db)
	}
	return deremsmodels.GroupGatewayRights(
		qm.Where("group_id = ?", groupID),
		qm.Where("disabled_at IS NULL")).All(repo.db)
}

func (repo defaultUserRepository) GetWebpagesPermissionByGroupTypeID(groupTypeID int64) ([]*deremsmodels.GroupTypeWebpageRight, error) {
	return deremsmodels.GroupTypeWebpageRights(
		qm.Where("type_id = ?", groupTypeID)).All(repo.db)
}

func (repo defaultUserRepository) GetWebpageByWebpageID(webpagesID int64) (*deremsmodels.Webpage, error) {
	return deremsmodels.FindWebpage(repo.db, webpagesID)
}
