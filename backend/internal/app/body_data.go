package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateGroupBody godoc
type CreateGroupBody struct {
	Name string `form:"name" binding:"required,max=20"`
	// TypeID 3 is "Area maintainer" and 4 is "Field owner"
	TypeID   int `form:"typeID" binding:"required,oneof=3 4"`
	ParentID int `form:"parentID" binding:"required"`
}

// UpdateGroupBody godoc
type UpdateGroupBody struct {
	Name string `form:"name" binding:"required,max=20"`
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

// Validate godoc
func (b *UpdateGroupBody) Validate(c *gin.Context) (err error) {
	if err = c.BindJSON(b); err != nil {
		logrus.WithField("caused-by", err).Error()
	}
	return
}
