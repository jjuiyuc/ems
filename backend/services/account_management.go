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
	GetGroup(groupID int64) (getGroup *GetGroupResponse, err error)
	UpdateGroup(userID, groupID int64, body *app.UpdateGroupBody) (err error)
}

// GetGroupsResponse godoc
type GetGroupsResponse struct {
	Groups []GetGroupInfo `json:"groups"`
}

// GetGroupInfo godoc
type GetGroupInfo struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	TypeID   int64      `json:"typeID"`
	ParentID null.Int64 `json:"parentID"`
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
	user, err := s.repo.User.GetUserByUserID(userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	groups, err := s.repo.User.GetGroupsByGroupID(user.GroupID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetGroupsByGroupID",
			"err":       err,
		}).Error()
		return
	}

	var getGroupInfos []GetGroupInfo
	for _, group := range groups {
		getGroupInfo := GetGroupInfo{
			ID:       group.ID,
			Name:     group.Name,
			TypeID:   group.TypeID,
			ParentID: group.ParentID,
		}
		getGroupInfos = append(getGroupInfos, getGroupInfo)
	}
	getGroups = &GetGroupsResponse{
		Groups: getGroupInfos,
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

func (s defaultAccountManagementService) GetGroup(groupID int64) (getGroup *GetGroupResponse, err error) {
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
	if s.isOwnAccountGroup(userID, groupID) {
		err = e.ErrNewOwnAccountGroupUpdatedNotAllow
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
