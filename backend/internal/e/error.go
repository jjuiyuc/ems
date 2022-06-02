package e

import (
	"fmt"
	"time"
)

func NewKeyNotExistError(key string) error {
	return fmt.Errorf("Key %s does not exist", key)
}

func NewUserExpirationError(expirationDate time.Time) error {
	return fmt.Errorf("User is expired on %s, please contact admin to extend", expirationDate)
}

func NewUserLockedError(passwordLockCount int) error {
	return fmt.Errorf("PasswordRetryCount: over %d tries, please contact admin to unlock", passwordLockCount)
}
