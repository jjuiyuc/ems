package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateGroupBody godoc
type CreateGroupBody struct {
	Name     string `form:"name" binding:"required,max=20"`
	TypeID   int    `form:"typeID" binding:"required,oneof=3 4"`
	ParentID int    `form:"parentID" binding:"required"`
}

// BodyData godoc
type BodyData interface {
	validate(c *gin.Context) error
}

// Validate godoc
func (b *CreateGroupBody) Validate(c *gin.Context) (err error) {
	if err = c.BindJSON(b); err != nil {
		logrus.WithField("caused-by", err).Error()
	}
	return
}
