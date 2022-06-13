package services

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"

	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// UserService ...
type UserService interface {
	CreatePasswordToken(username string) (name, token string, err error)
	PasswordResetByPasswordToken(token, newPassword string) (err error)
	GetProfile(userID int) (user *deremsmodels.User, err error)
}

type defaultUserService struct {
	repo *repository.Repository
}

// NewUserService ...
func NewUserService(repo *repository.Repository) UserService {
	return &defaultUserService{repo}
}

// CreatePasswordToken ...
func (s defaultUserService) CreatePasswordToken(username string) (name, token string, err error) {
	user, err := s.repo.User.GetUserByUsername(username)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.GetUserByUsername",
			"err":       err,
		}).Error()
		return
	}

	token = uuid.New().String()
	user.ResetPWDToken = null.NewString(token, true)
	user.PWDTokenExpiry = null.NewTime(time.Now().Add(1*time.Hour), true)
	err = s.repo.User.UpdateUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.UpdateUser",
			"err":       err,
		}).Error()
		return
	}

	name = user.Name.String
	return
}

// PasswordResetByPasswordToken ...
func (s defaultUserService) PasswordResetByPasswordToken(token, newPassword string) (err error) {
	user, err := s.repo.User.GetUserByPasswordToken(token)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByPasswordToken",
			"err":       err,
		}).Error()
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "bcrypt.GenerateFromPassword",
			"err":       err,
		}).Error()
		return
	}
	user.Password = string(hashPassword[:])
	user.PasswordLastChanged = null.NewTime(time.Now(), true)
	user.ResetPWDToken = null.NewString("", true)
	err = s.repo.User.UpdateUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.UpdateUser",
			"err":       err,
		}).Error()
	}
	return
}

// GetProfile ...
func (s defaultUserService) GetProfile(userID int) (user *deremsmodels.User, err error) {
	user, err = s.repo.User.GetProfileByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.GetProfileByUserID",
			"err":       err,
		}).Error()
	}
	return
}
