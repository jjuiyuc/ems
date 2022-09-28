package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/e"
)

// StartTimeQuery godoc
type StartTimeQuery struct {
	StartTime time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
}

// PeriodQuery godoc
type PeriodQuery struct {
	StartTime time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
	EndTime   time.Time `form:"endTime" binding:"required,gtfield=StartTime" example:"UTC time in ISO-8601" format:"date-time"`
}

// ZoomableQuery godoc
type ZoomableQuery struct {
	Resolution string    `form:"resolution" binding:"required" enums:"hour"`
	StartTime  time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
	EndTime    time.Time `form:"endTime" binding:"required,gtfield=StartTime" example:"UTC time in ISO-8601" format:"date-time"`
}

// ResolutionWithPeriodQuery godoc
type ResolutionWithPeriodQuery struct {
	Resolution string    `form:"resolution" binding:"required" enums:"day,month"`
	StartTime  time.Time `form:"startTime" binding:"required" example:"UTC time in ISO-8601" format:"date-time"`
	EndTime    time.Time `form:"endTime" binding:"required,gtfield=StartTime" example:"UTC time in ISO-8601" format:"date-time"`
}

// Param godoc
type Param interface {
	validate(c *gin.Context) error
}

// StartTimeParam godoc
type StartTimeParam struct {
	GatewayUUID string
	Query       StartTimeQuery
}

// PeriodParam godoc
type PeriodParam struct {
	GatewayUUID string
	Query       PeriodQuery
}

// ZoomableParam godoc
type ZoomableParam struct {
	GatewayUUID string
	Query       ZoomableQuery
}

// ResolutionWithPeriodParam godoc
type ResolutionWithPeriodParam struct {
	GatewayUUID string
	Query       ResolutionWithPeriodQuery
}

func (p *StartTimeParam) validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

func (p *PeriodParam) validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

func (p *ZoomableParam) validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
		return
	}
	// TODO: Only supports hour now
	if p.Query.Resolution != "hour" {
		err = e.ErrNewUnexpectedResolution
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

func (p *ResolutionWithPeriodParam) validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
		return
	}
	if p.Query.Resolution != "day" && p.Query.Resolution != "month" {
		err = e.ErrNewUnexpectedResolution
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}
