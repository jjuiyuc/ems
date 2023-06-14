package services

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// AccountManagementService godoc
type AccountManagementService interface {
	GetGroups(userID int64) (getGroups *GetGroupsResponse, err error)
	CreateGroup(body *app.CreateGroupBody) (err error)
	GetGroup(userID, groupID int64) (getGroup *GetGroupResponse, err error)
	GetGroupGateways(groupID int64) (groupGateways []GroupGatewayInfo)
	UpdateGroup(executedUserID, groupID int64, body *app.UpdateGroupBody) (err error)
	DeleteGroup(executedUserID, groupID int64) (err error)
	GetUsers(userID int64) (getUsers *GetUsersResponse, err error)
	CreateUser(userID int64, body *app.CreateUserBody) error
	UpdateUser(executedUserID, userID int64, body *app.UpdateUserBody) (err error)
	DeleteUser(executedUserID, userID int64) (err error)
}

// GetGroupsResponse godoc
type GetGroupsResponse struct {
	Groups     []GetGroupInfo     `json:"groups"`
	GroupTypes []GetGroupTypeInfo `json:"groupTypes"`
}

// GetGroupInfo godoc
type GetGroupInfo struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	TypeID   int64      `json:"typeID"`
	ParentID null.Int64 `json:"parentID"`
}

// GetGroupTypeInfo godoc
type GetGroupTypeInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetGroupResponse godoc
type GetGroupResponse struct {
	Name     string             `json:"name"`
	TypeID   int64              `json:"typeID"`
	ParentID null.Int64         `json:"parentID"`
	Gateways []GroupGatewayInfo `json:"gateways"`
}

// GroupGatewayInfo godoc
type GroupGatewayInfo struct {
	GatewayID    string `json:"gatewayID"`
	LocationName string `json:"locationName"`
}

// GetUsersResponse godoc
type GetUsersResponse struct {
	Users []*repository.UserWrap `json:"users"`
}

type defaultAccountManagementService struct {
	repo *repository.Repository
}

// NewAccountManagementService godoc
func NewAccountManagementService(repo *repository.Repository) AccountManagementService {
	return &defaultAccountManagementService{repo}
}

func (s defaultAccountManagementService) GetGroups(userID int64) (getGroups *GetGroupsResponse, err error) {
	groups, err := s.repo.User.GetGroupsByUserID(nil, userID)
	if err != nil {
		return
	}
	groupTypes, err := s.repo.User.GetGroupTypes()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupTypes",
			"err":       err,
		}).Error()
		return
	}

	var (
		getGroupInfos     []GetGroupInfo
		getGroupTypeInfos []GetGroupTypeInfo
	)
	for _, group := range groups {
		getGroupInfo := GetGroupInfo{
			ID:       group.ID,
			Name:     group.Name,
			TypeID:   group.TypeID,
			ParentID: group.ParentID,
		}
		getGroupInfos = append(getGroupInfos, getGroupInfo)
	}
	for _, groupType := range groupTypes {
		getGroupTypeInfo := GetGroupTypeInfo{
			ID:   groupType.ID,
			Name: groupType.Name,
		}
		getGroupTypeInfos = append(getGroupTypeInfos, getGroupTypeInfo)
	}
	getGroups = &GetGroupsResponse{
		Groups:     getGroupInfos,
		GroupTypes: getGroupTypeInfos,
	}
	return
}

func (s defaultAccountManagementService) CreateGroup(body *app.CreateGroupBody) (err error) {
	if !s.validateGroupType(int64(body.ParentID), 2) {
		logrus.WithField("parent-id", int64(body.ParentID)).Error("validate-group-type-failed")
		err = e.ErrNewAccountParentGroupTypeUnexpected
		return
	}

	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	group := &deremsmodels.Group{
		Name:     body.Name,
		TypeID:   int64(body.TypeID),
		ParentID: null.Int64From(int64(body.ParentID)),
	}
	if s.repo.User.IsGroupNameExistedOnSameLevel(tx, group) {
		err = e.ErrNewAccountGroupNameOnSameLevelExist
		tx.Rollback()
		return
	}
	if err = s.repo.User.CreateGroup(tx, group); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.CreateGroup",
			"err":       err,
			"body":      *body,
		}).Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) validateGroupType(groupID, expectedTypeID int64) bool {
	group, err := s.repo.User.GetGroupByGroupID(nil, groupID)
	if err == nil && group.TypeID == expectedTypeID {
		return true
	}
	return false
}

func (s defaultAccountManagementService) GetGroup(userID, groupID int64) (getGroup *GetGroupResponse, err error) {
	if !s.authorizeGroupID(nil, userID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	group, err := s.repo.User.GetGroupByGroupID(nil, groupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupByGroupID",
			"err":       err,
		}).Error()
		return
	}
	gateways := s.GetGroupGateways(groupID)

	getGroup = &GetGroupResponse{
		Name:     group.Name,
		TypeID:   group.TypeID,
		ParentID: group.ParentID,
		Gateways: gateways,
	}
	return
}

func (s defaultAccountManagementService) GetGroupGateways(groupID int64) (groupGateways []GroupGatewayInfo) {
	gatewaysPermission, err := s.repo.User.GetGatewaysPermissionByGroupID(nil, groupID, false)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGatewaysPermissionByGroupID",
			"err":       err,
		}).Warn()
		return
	}
	for _, gatewayPermission := range gatewaysPermission {
		var groupGatewayInfo GroupGatewayInfo
		gateway, err := s.repo.Gateway.GetGatewayByGatewayID(gatewayPermission.GWID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"caused-by": "s.repo.Gateway.GetGatewayByGatewayID",
				"err":       err,
			}).Warn()
			continue
		}
		groupGatewayInfo.GatewayID = gateway.UUID

		location, err := s.repo.Location.GetLocationByLocationID(gatewayPermission.LocationID.Int64)
		if err == nil {
			groupGatewayInfo.LocationName = location.Name
		} else {
			logrus.WithFields(logrus.Fields{
				"caused-by": "s.repo.Location.GetLocationByLocationID",
				"err":       err,
			}).Warn()
		}

		groupGateways = append(groupGateways, groupGatewayInfo)
	}
	return
}

func (s defaultAccountManagementService) UpdateGroup(executedUserID, groupID int64, body *app.UpdateGroupBody) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if !s.authorizeGroupID(tx, executedUserID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}
	if s.isOwnAccountGroup(tx, executedUserID, groupID) {
		err = e.ErrNewOwnAccountGroupModifiedNotAllow
		logrus.WithField("caused-by", err).Error()
		tx.Rollback()
		return
	}

	group, err := s.repo.User.GetGroupByGroupID(tx, groupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupByGroupID",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	group.Name = body.Name
	if s.repo.User.IsGroupNameExistedOnSameLevel(tx, group) {
		err = e.ErrNewAccountGroupNameOnSameLevelExist
		tx.Rollback()
		return
	}
	if err = s.repo.User.UpdateGroup(tx, group); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.UpdateGroup",
			"err":       err,
			"body":      *body,
		}).Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) isOwnAccountGroup(tx *sql.Tx, userID, groupID int64) bool {
	user, err := s.repo.User.GetUserByUserID(tx, userID)
	if err == nil && user.GroupID == groupID {
		return true
	}
	return false
}

func (s defaultAccountManagementService) DeleteGroup(executedUserID, groupID int64) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if !s.authorizeGroupID(tx, executedUserID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}
	if err = s.checkDeletedRules(tx, executedUserID, groupID); err != nil {
		tx.Rollback()
		return
	}

	if err = s.repo.User.DeleteGroup(tx, executedUserID, groupID); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.DeleteGroup",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) checkDeletedRules(tx *sql.Tx, executedUserID, groupID int64) (err error) {
	if s.isOwnAccountGroup(tx, executedUserID, groupID) {
		err = e.ErrNewOwnAccountGroupModifiedNotAllow
		logrus.WithField("caused-by", err).Error()
		return
	}
	if s.repo.User.IsSubGroupExisted(tx, groupID) {
		err = e.ErrNewAccountGroupHasSubGroup
		logrus.WithField("caused-by", err).Error()
		return
	}
	if s.repo.User.IsUserExistedInGroup(tx, groupID) {
		err = e.ErrNewAccountGroupHasUser
		logrus.WithField("caused-by", err).Error()
	}
	return
}

func (s defaultAccountManagementService) authorizeGroupID(tx *sql.Tx, executedUserID, groupID int64) (exist bool) {
	if exist = s.repo.User.AuthorizeGroupID(tx, executedUserID, groupID); !exist {
		logrus.WithField("executedUserID", executedUserID).Error("authorize-group-id-failed")
	}
	return
}

func (s defaultAccountManagementService) GetUsers(userID int64) (getUsers *GetUsersResponse, err error) {
	groups, err := s.repo.User.GetGroupsByUserID(nil, userID)
	if err != nil {
		return
	}

	var groupIDs []interface{}
	for _, group := range groups {
		groupIDs = append(groupIDs, group.ID)
	}
	users, err := s.repo.User.GetUserWrapsByGroupIDs(groupIDs)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserWrapsByGroupIDs",
			"err":       err,
		}).Error()
		return
	}
	getUsers = &GetUsersResponse{
		Users: users,
	}
	return
}

func (s defaultAccountManagementService) CreateUser(userID int64, body *app.CreateUserBody) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if !s.authorizeGroupID(tx, userID, int64(body.GroupID)) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}
	if s.repo.User.IsUsernameExisted(tx, body.Username) {
		err = e.ErrNewAccountUsernameExist
		logrus.WithField("caused-by", err).Error()
		tx.Rollback()
		return
	}

	hashPassword, err := utils.CreateHashedPassword(body.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "utils.CreateHashedPassword",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	user := &deremsmodels.User{
		Username: body.Username,
		Password: hashPassword,
		Name:     null.StringFrom(body.Name),
		GroupID:  int64(body.GroupID),
	}
	if err = s.repo.User.CreateUser(tx, user); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.CreateUser",
			"err":       err,
			"body":      *body,
		}).Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) UpdateUser(executedUserID, userID int64, body *app.UpdateUserBody) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if !s.authorizeUserID(tx, executedUserID, userID) ||
		(body.GroupID > 0 && !s.authorizeGroupID(tx, executedUserID, int64(body.GroupID))) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}

	user, err := s.repo.User.GetUserByUserID(tx, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}
	if err = s.processUpdateUser(tx, user, body); err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) processUpdateUser(tx *sql.Tx, user *deremsmodels.User, body *app.UpdateUserBody) (err error) {
	if body.Password != "" {
		var hashPassword string
		hashPassword, err = utils.CreateHashedPassword(body.Password)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"caused-by": "utils.CreateHashedPassword",
				"err":       err,
			}).Error()
			return
		}
		user.Password = hashPassword
	}
	if body.Name != "" {
		user.Name = null.StringFrom(body.Name)
	}
	if body.GroupID > 0 {
		user.GroupID = int64(body.GroupID)
	}
	if body.Unlock {
		user.PasswordRetryCount = null.IntFrom(0)
		user.LockedAt = null.TimeFromPtr(nil)
	}
	if err = s.repo.User.UpdateUser(tx, user); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
			"body":      *body,
		}).Error()
	}
	return
}

func (s defaultAccountManagementService) DeleteUser(executedUserID, userID int64) (err error) {
	tx, err := models.GetDB().BeginTx(context.Background(), nil)
	if err != nil {
		return
	}

	if !s.authorizeUserID(tx, executedUserID, userID) {
		err = e.ErrNewAuthPermissionNotAllow
		tx.Rollback()
		return
	}
	if executedUserID == userID {
		err = e.ErrNewOwnAccountDeletedNotAllow
		logrus.WithField("caused-by", err).Error()
		tx.Rollback()
		return
	}

	if err = s.repo.User.DeleteUser(tx, executedUserID, userID); err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.DeleteUser",
			"err":       err,
		}).Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (s defaultAccountManagementService) authorizeUserID(tx *sql.Tx, executedUserID, userID int64) (exist bool) {
	if exist = s.repo.User.AuthorizeUserID(tx, executedUserID, userID); !exist {
		logrus.WithField("executedUserID", executedUserID).Error("authorize-user-id-failed")
	}
	return
}
