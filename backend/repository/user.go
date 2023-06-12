package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
)

// UserWrap godoc
type UserWrap struct {
	ID              int64       `json:"id"`
	Username        string      `json:"username"`
	Name            null.String `json:"name"`
	LockedAt        null.Time   `json:"lockedAt"`
	GroupID         int64       `json:"groupID"`
	GroupName       string      `json:"groupName"`
	GroupParentID   null.Int64  `json:"groupParentID"`
	GroupParentName null.String `json:"groupParentName"`
}

// UserRepository godoc
type UserRepository interface {
	InsertLoginLog(loginLog *deremsmodels.LoginLog) error
	CreateUser(tx *sql.Tx, user *deremsmodels.User) error
	UpdateUser(tx *sql.Tx, user *deremsmodels.User) (err error)
	DeleteUser(tx *sql.Tx, executedUserID, userID int64) (err error)
	GetLoginLogCount() (int64, error)
	GetUserByUserID(tx *sql.Tx, userID int64) (*deremsmodels.User, error)
	GetUserByUsername(username string) (*deremsmodels.User, error)
	GetUserByPasswordToken(token string) (*deremsmodels.User, error)
	GetUserWrapsByGroupIDs(groupIDs []interface{}) ([]*UserWrap, error)
	IsUserExistedInGroup(tx *sql.Tx, groupID int64) bool
	IsUsernameExisted(tx *sql.Tx, username string) bool
	CreateGroup(tx *sql.Tx, group *deremsmodels.Group) (err error)
	UpdateGroup(tx *sql.Tx, group *deremsmodels.Group) (err error)
	DeleteGroup(tx *sql.Tx, executedUserID, groupID int64) (err error)
	IsGroupNameExistedOnSameLevel(tx *sql.Tx, group *deremsmodels.Group) bool
	GetGroupByGroupID(tx *sql.Tx, groupID int64) (*deremsmodels.Group, error)
	GetGroupsByGroupID(tx *sql.Tx, groupID int64) ([]*deremsmodels.Group, error)
	GetGroupsByUserID(tx *sql.Tx, userID int64) ([]*deremsmodels.Group, error)
	AuthorizeGroupID(tx *sql.Tx, executedUserID, groupID int64) (exist bool)
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

func (repo defaultUserRepository) CreateUser(tx *sql.Tx, user *deremsmodels.User) error {
	return user.Insert(repo.getExecutor(tx), boil.Infer())
}

// UpdateUser godoc
func (repo defaultUserRepository) UpdateUser(tx *sql.Tx, user *deremsmodels.User) (err error) {
	user.UpdatedAt = time.Now().UTC()
	_, err = user.Update(repo.getExecutor(tx), boil.Infer())
	return
}

func (repo defaultUserRepository) DeleteUser(tx *sql.Tx, executedUserID, userID int64) (err error) {
	exec := repo.getExecutor(tx)
	user, err := deremsmodels.FindUser(exec, userID)
	if err == nil {
		now := time.Now().UTC()
		user.UpdatedAt = now
		user.DeletedAt = null.TimeFrom(now)
		user.DeletedBy = null.Int64From(executedUserID)
		_, err = user.Update(exec, boil.Infer())
	}
	return
}

// GetLoginLogCount godoc
func (repo defaultUserRepository) GetLoginLogCount() (int64, error) {
	return deremsmodels.LoginLogs().Count(repo.db)
}

// GetUserByUserID godoc
func (repo defaultUserRepository) GetUserByUserID(tx *sql.Tx, userID int64) (*deremsmodels.User, error) {
	return deremsmodels.FindUser(repo.getExecutor(tx), userID)
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

func (repo defaultUserRepository) GetUserWrapsByGroupIDs(groupIDs []interface{}) (users []*UserWrap, err error) {
	users = make([]*UserWrap, 0)
	err = deremsmodels.NewQuery(
		qm.Select(
			"u.id AS id",
			"u.username AS username",
			"u.name AS name",
			"u.locked_at AS locked_at",
			"u.group_id AS group_id",
			"g.name AS group_name",
			"g.parent_id AS group_parent_id",
			"g2.name AS group_parent_name",
		),
		qm.From("user AS u"),
		qm.InnerJoin("`group` AS g ON u.group_id = g.id"),
		qm.LeftOuterJoin("`group` AS g2 ON g.parent_id = g2.id"),
		qm.WhereIn("u.group_id IN ?", groupIDs...),
		qm.Where("u.deleted_at IS NULL"),
		qm.OrderBy("u.id"),
	).Bind(context.Background(), models.GetDB(), &users)
	return
}

func (repo defaultUserRepository) IsUserExistedInGroup(tx *sql.Tx, groupID int64) (exist bool) {
	exist, _ = deremsmodels.Users(
		qm.Where("group_id = ?", groupID),
		qm.Where("deleted_at IS NULL")).Exists(repo.getExecutor(tx))
	return
}

func (repo defaultUserRepository) IsUsernameExisted(tx *sql.Tx, username string) (exist bool) {
	exist, _ = deremsmodels.Users(
		qm.Where("username = ?", username),
		qm.Where("deleted_at IS NULL")).Exists(repo.getExecutor(tx))
	return
}

func (repo defaultUserRepository) CreateGroup(tx *sql.Tx, group *deremsmodels.Group) (err error) {
	return group.Insert(repo.getExecutor(tx), boil.Infer())
}

func (repo defaultUserRepository) UpdateGroup(tx *sql.Tx, group *deremsmodels.Group) (err error) {
	group.UpdatedAt = time.Now().UTC()
	_, err = group.Update(repo.getExecutor(tx), boil.Infer())
	return
}

func (repo defaultUserRepository) DeleteGroup(tx *sql.Tx, executedUserID, groupID int64) (err error) {
	exec := repo.getExecutor(tx)
	group, err := deremsmodels.FindGroup(exec, groupID)
	if err == nil {
		now := time.Now().UTC()
		group.UpdatedAt = now
		group.DeletedAt = null.TimeFrom(now)
		group.DeletedBy = null.Int64From(executedUserID)
		_, err = group.Update(exec, boil.Infer())
	}
	return
}

func (repo defaultUserRepository) IsGroupNameExistedOnSameLevel(tx *sql.Tx, group *deremsmodels.Group) (exist bool) {
	exist, _ = deremsmodels.Groups(
		qm.Where("name = ?", group.Name),
		qm.Where("parent_id = ?", group.ParentID),
		qm.Where("deleted_at IS NULL")).Exists(repo.getExecutor(tx))
	return
}

func (repo defaultUserRepository) GetGroupByGroupID(tx *sql.Tx, groupID int64) (*deremsmodels.Group, error) {
	return deremsmodels.FindGroup(repo.getExecutor(tx), groupID)
}

func (repo defaultUserRepository) GetGroupsByGroupID(tx *sql.Tx, groupID int64) ([]*deremsmodels.Group, error) {
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
		SELECT * FROM group_path;`, "`group`", "`group`"), groupID)).All(repo.getExecutor(tx))
}

func (repo defaultUserRepository) GetGroupsByUserID(tx *sql.Tx, userID int64) ([]*deremsmodels.Group, error) {
	return deremsmodels.Groups(
		qm.SQL(fmt.Sprintf(`
		WITH RECURSIVE group_path AS
		(
		SELECT *
			FROM %s
			WHERE id = (
				SELECT group_id
				FROM user
				WHERE id = ?
			)
		UNION ALL
		SELECT g.*
			FROM group_path AS gp JOIN %s AS g
			ON gp.id = g.parent_id
			AND g.deleted_at IS NULL
		)
		SELECT * FROM group_path;`, "`group`", "`group`"), userID)).All(repo.getExecutor(tx))
}

func (repo defaultUserRepository) AuthorizeGroupID(tx *sql.Tx, executedUserID, groupID int64) (exist bool) {
	exist, _ = deremsmodels.Groups(
		qm.SQL(fmt.Sprintf(`
		WITH RECURSIVE group_path AS
		(
		SELECT *
			FROM %s
			WHERE id = (
				SELECT group_id
				FROM user
				WHERE id = ?
			)
		UNION ALL
		SELECT g.*
			FROM group_path AS gp JOIN %s AS g
			ON gp.id = g.parent_id
			AND g.deleted_at IS NULL
		)
		SELECT id FROM group_path WHERE id = ?;`, "`group`", "`group`"), executedUserID, groupID)).Exists(repo.getExecutor(tx))
	return
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

func (repo defaultUserRepository) getExecutor(tx *sql.Tx) boil.Executor {
	if tx == nil {
		return repo.db
	}
	return tx
}
