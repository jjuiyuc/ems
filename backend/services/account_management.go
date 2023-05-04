package services

import (
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/repository"
)

// AccountManagementService godoc
type AccountManagementService interface {
	GetSubGroups(userID int64) (subGroups *SubGroupsResponse, err error)
}

// SubGroupsResponse godoc
type SubGroupsResponse struct {
	Groups []SubGroupInfo `json:"groups"`
}

// SubGroupInfo godoc
type SubGroupInfo struct {
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

func (s defaultAccountManagementService) GetSubGroups(userID int64) (subGroups *SubGroupsResponse, err error) {
	user, err := s.repo.User.GetUserByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	groups, err := s.repo.User.GetSubGroupsByGroupID(user.GroupID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetSubGroupsByGroupID",
			"err":       err,
		}).Error()
		return
	}

	var subGroupInfos []SubGroupInfo
	for _, group := range groups {
		subGroupInfo := SubGroupInfo{
			ID:       group.ID,
			Name:     group.Name,
			TypeID:   group.TypeID,
			ParentID: group.ParentID,
		}
		subGroupInfos = append(subGroupInfos, subGroupInfo)
	}
	subGroups = &SubGroupsResponse{
		Groups: subGroupInfos,
	}
	return
}
