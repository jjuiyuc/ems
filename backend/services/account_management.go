package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/app"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// AccountManagementService godoc
type AccountManagementService interface {
	GetGroups(userID int64) (getGroups *GetGroupsResponse, err error)
	CreateGroup(body *app.CreateGroupBody) (errCode int, err error)
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
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	groups, err := s.repo.User.GetGroupsByGroupID(user.GroupID)
	if err != nil {
		log.WithFields(log.Fields{
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

func (s defaultAccountManagementService) CreateGroup(body *app.CreateGroupBody) (errCode int, err error) {
	group := &deremsmodels.Group{
		Name:     body.Name,
		TypeID:   int64(body.TypeID),
		ParentID: null.Int64From(int64(body.ParentID)),
	}

	errCode, err = s.repo.User.CreateGroup(group)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.CreateGroup",
			"err":       err,
			"body":      *body,
		}).Error()
	}
	return
}
