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

// UserService godoc
type UserService interface {
	CreatePasswordToken(username string) (name, token string, err error)
	PasswordResetByPasswordToken(token, newPassword string) (err error)
	GetProfile(userID int64) (profile *ProfileResponse, err error)
}

type defaultUserService struct {
	repo *repository.Repository
}

// ProfileResponse godoc
type ProfileResponse struct {
	*deremsmodels.User
	Gateways []GatewayInfo `json:"gateways"`
}

// GatewayInfo godoc
type GatewayInfo struct {
	GatewayID string `json:"gatewayID"`
	Address   string `json:"address"`
}

// NewUserService godoc
func NewUserService(repo *repository.Repository) UserService {
	return &defaultUserService{repo}
}

// CreatePasswordToken godoc
func (s defaultUserService) CreatePasswordToken(username string) (name, token string, err error) {
	user, err := s.repo.User.GetUserByUsername(username)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUsername",
			"err":       err,
		}).Error()
		return
	}

	token = uuid.New().String()
	user.ResetPWDToken = null.NewString(token, true)
	user.PWDTokenExpiry = null.NewTime(time.Now().UTC().Add(1*time.Hour), true)
	err = s.repo.User.UpdateUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
		return
	}

	name = user.Name.String
	return
}

// PasswordResetByPasswordToken godoc
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
	user.PasswordLastChanged = null.NewTime(time.Now().UTC(), true)
	user.ResetPWDToken = null.NewString("", true)
	err = s.repo.User.UpdateUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
	}
	return
}

// GetProfile godoc
func (s defaultUserService) GetProfile(userID int64) (profile *ProfileResponse, err error) {
	// Get user information
	user, err := s.repo.User.GetUserByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}

	// Get gateway information
	gateways, err := s.repo.Gateway.GetGatewaysByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.Gateway.GetGatewaysByUserID",
			"err":       err,
		}).Error()
		return
	}
	gatewayInfos := []GatewayInfo{}
	for _, gateway := range gateways {
		log.Debug("gateway.UUID: ", gateway.UUID)
		location, err := s.repo.Location.GetLocationByLocationID(gateway.LocationID)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "s.repo.Location.GetLocationByLocationID",
				"err":       err,
			}).Error()
			continue
		}
		log.Debug("location.Address: ", location.Address)

		gatewayInfo := GatewayInfo{
			GatewayID: gateway.UUID,
			Address:   location.Address.String,
		}
		gatewayInfos = append(gatewayInfos, gatewayInfo)
	}

	profile = &ProfileResponse{
		User:     user,
		Gateways: gatewayInfos,
	}
	return
}
