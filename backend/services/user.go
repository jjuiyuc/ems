package services

import (
	log "github.com/sirupsen/logrus"

	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// UserService ...
type UserService struct {
	repo repository.UserRepository
}

// NewUserService ...
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

// GetProfile ...
func (s *UserService) GetProfile(userID int) (user *deremsmodels.User, err error) {
	user, err = s.repo.GetProfileByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.GetProfileByUserID",
			"err":       err,
		}).Error()
	}
	return
}
