package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetGroupURI godoc
type GetGroupURI struct {
	GroupID int64 `uri:"groupid" binding:"required"`
}

// URI godoc
type URI interface {
	validate(c *gin.Context) error
}

// Validate godoc
func (u *GetGroupURI) Validate(c *gin.Context) (err error) {
	if err = c.ShouldBindUri(u); err != nil {
		logrus.WithField("caused-by", err).Error()
	}
	return
}
