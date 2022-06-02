package fixtures

import (
	"time"

	"github.com/volatiletech/null/v8"

	deremsmodels "der-ems/models/der-ems"
)

var UtUser = &deremsmodels.User{
	ID:             1,
	Username:       "ut-user@gmail.com",
	Password:       "testing123",
	ExpirationDate: null.NewTime(time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), true),
}
