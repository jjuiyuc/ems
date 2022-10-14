package fixtures

import (
	"time"

	"github.com/volatiletech/null/v8"

	deremsmodels "der-ems/models/der-ems"
)

// UtUser godoc
var UtUser = &deremsmodels.User{
	ID:             1,
	Username:       "ut-user@gmail.com",
	Password:       "testing123",
	ExpirationDate: null.NewTime(time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), true),
}

// UtCustomer godoc
var UtCustomer = &deremsmodels.Customer{
	ID:             1,
	CustomerNumber: "00001",
	FieldNumber:    "001",
	Address:        null.NewString("宜蘭縣五結鄉大吉五路157巷68號", true),
	Lat:            null.NewFloat64(24.70155508690467, true),
	Lng:            null.NewFloat64(121.7973398847259, true),
	WeatherLat:     null.NewFloat32(24.75, true),
	WeatherLng:     null.NewFloat32(121.75, true),
	Timezone:       null.NewString("+0800", true),
	TOULocationID:  null.NewInt64(1, true),
	VoltageType:    null.NewString("Low voltage", true),
	TOUType:        null.NewString("Two-section", true),
}

// UtGateway godoc
var UtGateway = &deremsmodels.Gateway{
	ID:         1,
	UUID:       "0E0BA27A8175AF978C49396BDE9D7A1E",
	CustomerID: 1,
}
