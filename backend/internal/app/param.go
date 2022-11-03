package app

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
	Resolution string    `form:"resolution" binding:"required" enums:"hour,5minute"`
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

// Validate godoc
func (p *StartTimeParam) Validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

// Validate godoc
func (p *PeriodParam) Validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

// Validate godoc
func (p *ZoomableParam) Validate(c *gin.Context) (err error) {
	p.GatewayUUID = c.Param("gwid")
	log.Debug("gatewayUUID: ", p.GatewayUUID)

	if err = c.BindQuery(&p.Query); err != nil {
		log.WithFields(log.Fields{"caused-by": err}).Error()
		return
	}
	if p.Query.Resolution != "hour" && p.Query.Resolution != "5minute" {
		err = e.ErrNewUnexpectedResolution
		log.WithFields(log.Fields{"caused-by": err}).Error()
	}
	return
}

// GetEndTimeIndex godoc
func (p *ZoomableParam) GetEndTimeIndex() (endTimeIndex time.Time) {
	switch p.Query.Resolution {
	case "hour":
		endTimeIndex = p.Query.StartTime.Add(1 * time.Hour)
	case "5minute":
		endTimeIndex = p.Query.StartTime.Add(5 * time.Minute)
	}
	return
}

// Validate godoc
func (p *ResolutionWithPeriodParam) Validate(c *gin.Context) (err error) {
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

// GetEndTimeIndex godoc
func (p *ResolutionWithPeriodParam) GetEndTimeIndex() (endTimeIndex time.Time) {
	switch p.Query.Resolution {
	case "day":
		endTimeIndex = p.Query.StartTime.AddDate(0, 0, 1)
	case "month":
		endTimeIndex = p.Query.StartTime.AddDate(0, 0, 1).AddDate(0, 1, 0).AddDate(0, 0, -1)
	}
	return
}
