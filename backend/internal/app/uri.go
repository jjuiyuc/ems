package app

// FieldURI godoc
type FieldURI struct {
	GatewayID string `uri:"gwid" binding:"required"`
}

// FieldDeviceURI godoc
type FieldDeviceURI struct {
	DeviceUUEID string `uri:"deviceuueid" binding:"required"`
}

// GatewayAndPeriodURI godoc
type GatewayAndPeriodURI struct {
	GatewayID string `uri:"gwid" binding:"required"`
	PeriodID  int64  `uri:"periodid" binding:"required"`
}

// GroupURI godoc
type GroupURI struct {
	GroupID int64 `uri:"groupid" binding:"required"`
}

// UserURI godoc
type UserURI struct {
	UserID int64 `uri:"userid" binding:"required"`
}
