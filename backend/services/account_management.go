package services

import (
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/app"
	"der-ems/internal/e"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// AccountManagementService godoc
type AccountManagementService interface {
	GetGroups(userID int64) (getGroups *GetGroupsResponse, err error)
	CreateGroup(body *app.CreateGroupBody) (err error)
	GetGroup(userID, groupID int64) (getGroup *GetGroupResponse, err error)
	UpdateGroup(userID, groupID int64, body *app.UpdateGroupBody) (err error)
	DeleteGroup(userID, groupID int64) (err error)
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

type defaultAccountManagementService struct {
	repo *repository.Repository
}

// NewAccountManagementService godoc
func NewAccountManagementService(repo *repository.Repository) AccountManagementService {
	return &defaultAccountManagementService{repo}
}

func (s defaultAccountManagementService) GetGroups(userID int64) (getGroups *GetGroupsResponse, err error) {
	groups, err := s.getGroupTreeNodes(userID)
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

func (s defaultAccountManagementService) getGroupTreeNodes(userID int64) (groups []*deremsmodels.Group, err error) {
	user, err := s.repo.User.GetUserByUserID(userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	groups, err = s.repo.User.GetGroupsByGroupID(user.GroupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupsByGroupID",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultAccountManagementService) CreateGroup(body *app.CreateGroupBody) (err error) {
	if !s.validateGroupType(int64(body.ParentID), 2) {
		logrus.WithField("parent-id", int64(body.ParentID)).Error("validate-group-type-failed")
		err = e.ErrNewAccountParentGroupTypeUnexpected
		return
	}

	group := &deremsmodels.Group{
		Name:     body.Name,
		TypeID:   int64(body.TypeID),
		ParentID: null.Int64From(int64(body.ParentID)),
	}
	if s.repo.User.IsGroupNameExistedOnSameLevel(group) {
		err = e.ErrNewAccountGroupNameOnSameLevelExist
		return
	}
	err = s.repo.User.CreateGroup(group)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.CreateGroup",
			"err":       err,
			"body":      *body,
		}).Error()
	}
	return
}

func (s defaultAccountManagementService) validateGroupType(groupID, expectedTypeID int64) bool {
	group, err := s.repo.User.GetGroupByGroupID(groupID)
	if err == nil && group.TypeID == expectedTypeID {
		return true
	}
	return false
}

func (s defaultAccountManagementService) GetGroup(userID, groupID int64) (getGroup *GetGroupResponse, err error) {
	if !s.authorizeGroupID(userID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	group, err := s.repo.User.GetGroupByGroupID(groupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupByGroupID",
			"err":       err,
		}).Error()
		return
	}
	gateways := s.getGroupGateways(groupID)

	getGroup = &GetGroupResponse{
		Name:     group.Name,
		TypeID:   group.TypeID,
		ParentID: group.ParentID,
		Gateways: gateways,
	}
	return
}

func (s defaultAccountManagementService) getGroupGateways(groupID int64) (groupGateways []GroupGatewayInfo) {
	gatewaysPermission, err := s.repo.User.GetGatewaysPermissionByGroupID(groupID, false)
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

func (s defaultAccountManagementService) UpdateGroup(userID, groupID int64, body *app.UpdateGroupBody) (err error) {
	if !s.authorizeGroupID(userID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	if s.isOwnAccountGroup(userID, groupID) {
		err = e.ErrNewOwnAccountGroupModifiedNotAllow
		logrus.WithField("caused-by", err).Error()
		return
	}

	group, err := s.repo.User.GetGroupByGroupID(groupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupByGroupID",
			"err":       err,
		}).Error()
		return
	}
	group.Name = body.Name
	if s.repo.User.IsGroupNameExistedOnSameLevel(group) {
		err = e.ErrNewAccountGroupNameOnSameLevelExist
		return
	}
	err = s.repo.User.UpdateGroup(group)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.UpdateGroup",
			"err":       err,
			"body":      *body,
		}).Error()
	}
	return
}

func (s defaultAccountManagementService) isOwnAccountGroup(userID, groupID int64) bool {
	user, err := s.repo.User.GetUserByUserID(userID)
	if err == nil && user.GroupID == groupID {
		return true
	}
	return false
}

func (s defaultAccountManagementService) DeleteGroup(userID, groupID int64) (err error) {
	if !s.authorizeGroupID(userID, groupID) {
		err = e.ErrNewAuthPermissionNotAllow
		return
	}
	err = s.checkDeletedRules(userID, groupID)
	if err != nil {
		return
	}

	err = s.repo.User.DeleteGroup(userID, groupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.DeleteGroup",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultAccountManagementService) checkDeletedRules(userID, groupID int64) (err error) {
	if s.isOwnAccountGroup(userID, groupID) {
		err = e.ErrNewOwnAccountGroupModifiedNotAllow
		logrus.WithField("caused-by", err).Error()
		return
	}
	if s.isSubGroupExisted(groupID) {
		err = e.ErrNewAccountGroupHasSubGroup
		logrus.WithField("caused-by", err).Error()
		return
	}
	if s.isUserExistedInGroup(groupID) {
		err = e.ErrNewAccountGroupHasUser
		logrus.WithField("caused-by", err).Error()
	}
	return
}

func (s defaultAccountManagementService) isSubGroupExisted(groupID int64) bool {
	groups, err := s.repo.User.GetGroupsByGroupID(groupID)
	if err == nil {
		for _, group := range groups {
			if group.ParentID.Int64 == groupID {
				return true
			}
		}
	}
	return false
}

func (s defaultAccountManagementService) isUserExistedInGroup(groupID int64) bool {
	count, _ := s.repo.User.GetUserCountByGroupID(groupID)
	return count > 0
}

func (s defaultAccountManagementService) authorizeGroupID(userID, groupID int64) bool {
	groups, err := s.getGroupTreeNodes(userID)
	if err == nil {
		for _, group := range groups {
			if group.ID == groupID {
				return true
			}
		}
	}
	logrus.WithField("userID", userID).Error("authorize-group-id-failed")
	return false
}
