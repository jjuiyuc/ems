package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GroupURI godoc
type GroupURI struct {
	GroupID int64 `uri:"groupid" binding:"required"`
}

// URI godoc
type URI interface {
	validate(c *gin.Context) error
}

// Validate godoc
func (u *GroupURI) Validate(c *gin.Context) (err error) {
	if err = c.ShouldBindUri(u); err != nil {
		logrus.WithField("caused-by", err).Error()
	}
	return
}
