package testdata

import (
	"time"

	"github.com/volatiletech/null/v8"

	deremsmodels "der-ems/models/der-ems"
)

// UtUser godoc
var UtUser = &deremsmodels.User{
	ID:             1,
	Username:       "ut-user@gmail.com",
	GroupID:        1,
	Password:       "testing123",
	ExpirationDate: null.TimeFrom(time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)),
	CreatedAt:      time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt:      time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
}

// UtLocation godoc
var UtLocation = &deremsmodels.Location{
	ID:            1,
	Name:          "Field A",
	Address:       null.StringFrom("宜蘭縣五結鄉大吉五路157巷68號"),
	Lat:           null.Float64From(24.70155508690467),
	Lng:           null.Float64From(121.7973398847259),
	WeatherLat:    null.Float32From(24.75),
	WeatherLng:    null.Float32From(121.75),
	TOULocationID: null.Int64From(1),
	VoltageType:   null.StringFrom("Low voltage"),
	TOUType:       null.StringFrom("Two-section"),
}

// UtGateway godoc
var UtGateway = &deremsmodels.Gateway{
	ID:         1,
	UUID:       "0E0BA27A8175AF978C49396BDE9D7A1E",
	LocationID: null.Int64From(1),
}
