package services

import (
	"github.com/sirupsen/logrus"

	"der-ems/repository"
)

// FieldManagementService godoc
type FieldManagementService interface {
	GetFields(userID int64) (getFields *GetFieldsResponse, err error)
}

// GetFieldsResponse godoc
type GetFieldsResponse struct {
	Gateways []GroupGatewayInfo `json:"gateways"`
}

type defaultFieldManagementService struct {
	repo *repository.Repository
	accountManagement AccountManagementService
}

// NewFieldManagementService godoc
func NewFieldManagementService(repo *repository.Repository, accountManagement AccountManagementService) FieldManagementService {
	return &defaultFieldManagementService{repo, accountManagement}
}

func (s defaultFieldManagementService) GetFields(userID int64) (getFields *GetFieldsResponse, err error) {
	user, err := s.repo.User.GetUserByUserID(nil, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}
	gateways := s.accountManagement.GetGroupGateways(user.GroupID)
	getFields = &GetFieldsResponse{
		Gateways: gateways,
	}
	return
}
