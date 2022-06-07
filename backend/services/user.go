package services

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

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

// PasswordLost ...
func (s *UserService) PasswordLost(c *gin.Context, username string) (err error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.GetUserByUsername",
			"err":       err,
		}).Error()
		return
	}

	// Create temporary password
	token := uuid.New().String()
	user.Password = token
	user.PasswordResetExpiry = null.NewTime(time.Now().Add(1*time.Hour), true)
	err = s.repo.UpdateUser(user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.UpdateUser",
			"err":       err,
		}).Error()
		return
	}

	// Send temporary password by email
	referrer, err := getReferrerBase(c)
	if err != nil {
		return
	}
	err = SendResetEmail(user.Name.String, user.Username, referrer, token)
	return
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

func getReferrerBase(c *gin.Context) (referrer string, err error) {
	u, err := url.Parse(c.Request.Referer())
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "getReferrerBase url.Parse",
			"err":       err,
		}).Error()
		return
	}
	referrer = u.Scheme + "://" + u.Host
	return
}
