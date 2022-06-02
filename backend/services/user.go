package services

import (
	log "github.com/sirupsen/logrus"

	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

func GetProfile(userID int) (user *deremsmodels.User, err error) {
	user, err = repository.GetProfileByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "repository.GetProfileByUserID",
			"err":       err,
		}).Error()
	}
	return
}
