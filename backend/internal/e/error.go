package e

import (
	"errors"
	"fmt"
	"time"
)

var (
	// ErrNewUnexpectedJSONInput godoc
	ErrNewUnexpectedJSONInput = errors.New("unexpected end of JSON input")
	// ErrNewUnexpectedTimeRange godoc
	ErrNewUnexpectedTimeRange = errors.New("unexpected start time and end time")
)

// ErrNewKeyNotExist godoc
func ErrNewKeyNotExist(key string) error {
	return fmt.Errorf("Key %s does not exist", key)
}

// ErrNewUserExpiration godoc
func ErrNewUserExpiration(expirationDate time.Time) error {
	return fmt.Errorf("User is expired on %s, please contact admin to extend", expirationDate)
}

// ErrNewUserLocked godoc
func ErrNewUserLocked(passwordLockCount int) error {
	return fmt.Errorf("PasswordRetryCount: over %d tries, please contact admin to unlock", passwordLockCount)
}
