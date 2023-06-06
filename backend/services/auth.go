package services

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

const passwordLockCount = 5

// AuthService godoc
type AuthService interface {
	Login(username, password string) (user *deremsmodels.User, groupType int64, errCode int, err error)
	CreateLoginLog(user *deremsmodels.User, token string) (err error)
}

type defaultAuthService struct {
	repo *repository.Repository
}

// NewAuthService godoc
func NewAuthService(repo *repository.Repository) AuthService {
	return &defaultAuthService{repo}
}

// Login godoc
func (s defaultAuthService) Login(username, password string) (user *deremsmodels.User, groupType int64, errCode int, err error) {
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
	now := time.Now().UTC()
	if user.ExpirationDate.Valid != false && user.ExpirationDate.Time.Before(now) {
		errCode = e.ErrAuthUserExpirated
		err = e.ErrNewUserExpiration(user.ExpirationDate.Time)
		log.WithFields(log.Fields{
			"caused-by": "user.ExpirationDate",
			"err":       err,
		}).Error()
		return
	}

	// Check password retry count
	if user.PasswordRetryCount.Int >= passwordLockCount {
		errCode = e.ErrAuthUserLocked
		err = e.ErrNewUserLocked(passwordLockCount)
		log.WithFields(log.Fields{
			"caused-by": "user.PasswordRetryCount",
			"err":       err,
		}).Error()
		return
	}

	// Check password
	nowPasswordRetryCount := user.PasswordRetryCount.Int
	if err = utils.ComparePassword(password, user.Password); err != nil {
		errCode = e.ErrAuthPasswordNotMatch
		user.PasswordRetryCount = null.IntFrom(nowPasswordRetryCount + 1)
		if user.PasswordRetryCount.Int == passwordLockCount {
			now := time.Now().UTC()
			user.LockedAt = null.TimeFrom(now)
		}
		s.repo.User.UpdateUser(user)
		return
	}
	if nowPasswordRetryCount > 0 {
		user.PasswordRetryCount = null.IntFrom(0)
		s.repo.User.UpdateUser(user)
	}

	group, err := s.repo.User.GetGroupByGroupID(nil, user.GroupID)
	if err != nil {
		errCode = e.ErrAuthUserNotExist
		log.WithFields(log.Fields{
			"caused-by": "repository.GetGroupByGroupID",
			"err":       err,
		}).Error()
		return
	}
	groupType = group.TypeID

	return
}

// CreateLoginLog godoc
func (s defaultAuthService) CreateLoginLog(user *deremsmodels.User, token string) (err error) {
	loginLog := &deremsmodels.LoginLog{
		UserID: null.Int64From(user.ID),
	}

	err = s.repo.User.InsertLoginLog(loginLog)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.InsertLoginLog",
			"err":       err,
		}).Error()
	}
	return
}
