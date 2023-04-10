package utils

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// ComparePassword godoc
func ComparePassword(rawPassword, hashedPassword string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "bcrypt.CompareHashAndPassword",
			"err":       err,
		}).Error()
	}
	return
}

// CreateHashedPassword godoc
func CreateHashedPassword(rawPassword string) (hashedPassword string, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "bcrypt.GenerateFromPassword",
			"err":       err,
		}).Error()
		return
	}
	hashedPassword = string(password[:])
	return
}
