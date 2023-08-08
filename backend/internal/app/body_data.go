package app

import (
	"time"
)

// EnableFieldBody godoc
type EnableFieldBody struct {
	Enable *bool `form:"enable" binding:"required"`
}

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

// CreateUserBody godoc
type CreateUserBody struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required,max=50"`
	Name     string `form:"name" binding:"required,max=20"`
	GroupID  int    `form:"groupID" binding:"required"`
}

// UpdateUserBody godoc
type UpdateUserBody struct {
	Password string `form:"password" binding:"max=50"`
	Name     string `form:"name" binding:"max=20"`
	GroupID  int    `form:"groupID"`
	Unlock   bool   `form:"unlock"`
}

// CreateFieldBody godoc
type CreateFieldBody struct {
	GatewayID    string            `form:"gatewayID" binding:"required"`
	LocationName string            `form:"locationName" binding:"required"`
	Address      string            `form:"address" binding:"required"`
	Lat          float64           `form:"lat" binding:"required,latitude"`
	Lng          float64           `form:"lng" binding:"required,longitude"`
	PowerCompany string            `form:"powerCompany" binding:"oneof='TPC'"`
	VoltageType  string            `form:"voltageType" binding:"oneof='Low voltage' 'High voltage'"`
	TOUType      string            `form:"touType" binding:"oneof='Two-section'"`
	Enable       *bool             `form:"enable" binding:"required"`
	Devices      []FieldDeviceInfo `form:"devices" binding:"required,dive"`
}

// FieldDeviceInfo godoc
type FieldDeviceInfo struct {
	ModelID       int64                 `form:"modelID" binding:"required"`
	ModbusID      int                   `form:"modbusID" binding:"required"`
	UUEID         string                `form:"uueid" binding:"required"`
	PowerCapacity float32               `form:"powerCapacity" binding:"required"`
	ExtraInfo     *FieldDeviceExtraInfo `form:"extraInfo"`
	SubDevices    []FieldSubDeviceInfo  `form:"subDevices" binding:"dive"`
}

// FieldSubDeviceInfo godoc
type FieldSubDeviceInfo struct {
	ModelID       int64                 `form:"modelID" binding:"required"`
	PowerCapacity float32               `form:"powerCapacity" binding:"required"`
	ExtraInfo     *FieldDeviceExtraInfo `form:"extraInfo"`
}

// FieldDeviceExtraInfo godoc
type FieldDeviceExtraInfo struct {
	ReservedForGridOutagePercent int     `form:"reservedForGridOutagePercent" binding:"required" json:"reservedForGridOutagePercent"`
	ChargingSources              string  `form:"chargingSources" binding:"required,oneof='Solar + Grid' 'Solar'" json:"chargingSources"`
	EnergyCapacity               float32 `form:"energyCapacity" binding:"required" json:"energyCapacity"`
	Voltage                      float32 `form:"voltage" binding:"required" json:"voltage"`
}

// UpdateFieldGroupsBody godoc
type UpdateFieldGroupsBody struct {
	Groups []FieldGroupInfo `form:"groups" binding:"required,dive"`
}

// FieldGroupInfo godoc
type FieldGroupInfo struct {
	ID    int64 `form:"id" binding:"required"`
	Check *bool `form:"check" binding:"required"`
}

// UpdateBatterySettingsBody godoc
type UpdateBatterySettingsBody struct {
	// According with luxpower logic, the sources need to be 'Solar + Grid' or 'Solar'
	ChargingSources              string `form:"chargingSources" binding:"required,oneof='Solar + Grid' 'Solar'"`
	ReservedForGridOutagePercent int    `form:"reservedForGridOutagePercent" binding:"required"`
}

// UpdateMeterSettingsBody godoc
type UpdateMeterSettingsBody struct {
	MaxDemandCapacity int `form:"maxDemandCapacity" binding:"required"`
}

// CreatePowerOutagePeriodsBody godoc
type CreatePowerOutagePeriodsBody struct {
	Periods []PowerOutagePeriodInfo `form:"periods" binding:"required,dive"`
}

// PowerOutagePeriodInfo godoc
type PowerOutagePeriodInfo struct {
	Type      string    `form:"type" binding:"required,oneof='advanceBlackout' 'evCharge'"`
	StartTime time.Time `form:"startTime" binding:"required,gt" format:"date-time"`
	EndTime   time.Time `form:"endTime" binding:"required,gtfield=StartTime" format:"date-time"`
}
