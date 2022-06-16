package services

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"

	"der-ems/internal/e"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

const passwordLockCount = 5

// AuthService ...
type AuthService interface {
	Login(username, password string) (user *deremsmodels.User, errCode int, err error)
	CreateLoginLog(user *deremsmodels.User, token string) (err error)
}

type defaultAuthService struct {
	repo *repository.Repository
}

// NewAuthService ...
func NewAuthService(repo *repository.Repository) AuthService {
	return &defaultAuthService{repo}
}

// Login ...
func (s defaultAuthService) Login(username, password string) (user *deremsmodels.User, errCode int, err error) {
	user, err = s.repo.User.GetUserByUsername(username)
	if err != nil {
		errCode = e.ErrAuthUserNotExist
		log.WithFields(log.Fields{
			"caused-by": "repository.GetUserByUsername",
			"err":       err,
		}).Error()
		return
	}

	// Check expiration date
	now := time.Now()
	if user.ExpirationDate.Valid != false && user.ExpirationDate.Time.Before(now) {
		errCode = e.ErrAuthUserExpirated
		err = e.NewUserExpirationError(user.ExpirationDate.Time)
		log.WithFields(log.Fields{
			"caused-by": "user.ExpirationDate",
			"err":       err,
		}).Error()
		return
	}

	// Check password retry count
	if user.PasswordRetryCount.Int >= passwordLockCount {
		errCode = e.ErrAuthUserLocked
		err = e.NewUserLockedError(passwordLockCount)
		log.WithFields(log.Fields{
			"caused-by": "user.PasswordRetryCount",
			"err":       err,
		}).Error()
		return
	}

	// Check password
	nowPasswordRetryCount := user.PasswordRetryCount.Int
	err = comparePassword(password, user.Password)
	if err != nil {
		errCode = e.ErrAuthPasswordNotMatch
		log.WithFields(log.Fields{
			"caused-by": "comparePassword",
			"err":       err,
		}).Error()

		user.PasswordRetryCount = null.NewInt(nowPasswordRetryCount+1, true)
		if user.PasswordRetryCount.Int == passwordLockCount {
			now := time.Now()
			user.LockedAt = null.NewTime(now, true)
		}
		s.repo.User.UpdateUser(user)
		return
	}
	if nowPasswordRetryCount > 0 {
		user.PasswordRetryCount = null.NewInt(0, true)
		s.repo.User.UpdateUser(user)
	}

	return
}

// CreateLoginLog ...
func (s defaultAuthService) CreateLoginLog(user *deremsmodels.User, token string) (err error) {
	loginLog := &deremsmodels.LoginLog{
		UserID: null.NewInt(user.ID, true),
	}

	err = s.repo.User.InsertLoginLog(loginLog)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.InsertLoginLog",
			"err":       err,
		}).Error()
	}
	return
}

func comparePassword(rawPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(rawPassword))
}
