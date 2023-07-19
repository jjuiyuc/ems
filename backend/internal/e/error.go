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
	// ErrNewMessageNotEqual godoc
	ErrNewMessageNotEqual = errors.New("message when not equal")
	// ErrNewMessageReceivedUnexpectedErr godoc
	ErrNewMessageReceivedUnexpectedErr = errors.New("message when received unexpected error")
	// ErrNewMessageGotNil godoc
	ErrNewMessageGotNil = errors.New("message when an error is expected but got nil")
	// ErrNewUnexpectedResolution godoc
	ErrNewUnexpectedResolution = errors.New("unexpected resolution")
	// ErrNewBillingsNotExist godoc
	ErrNewBillingsNotExist = errors.New("billings do not exist")
	// ErrNewAccountGroupNameOnSameLevelExist godoc
	ErrNewAccountGroupNameOnSameLevelExist = errors.New("account group name exists on the same level")
	// ErrNewAccountParentGroupTypeUnexpected godoc
	ErrNewAccountParentGroupTypeUnexpected = errors.New("account parent group type is unexpected")
	// ErrNewOwnAccountGroupModifiedNotAllow godoc
	ErrNewOwnAccountGroupModifiedNotAllow = errors.New("modify own account group does not allow")
	// ErrNewAccountGroupHasSubGroup godoc
	ErrNewAccountGroupHasSubGroup = errors.New("account group has sub group")
	// ErrNewAccountGroupHasUser godoc
	ErrNewAccountGroupHasUser = errors.New("account group has user")
	// ErrNewAuthPermissionNotAllow godoc
	ErrNewAuthPermissionNotAllow = errors.New("authorization permission does not allow")
	// ErrNewAccountUsernameExist godoc
	ErrNewAccountUsernameExist = errors.New("account username exists")
	// ErrNewOwnAccountDeletedNotAllow godoc
	ErrNewOwnAccountDeletedNotAllow = errors.New("delete own account does not allow")
	// ErrNewFieldIsDisabled godoc
	ErrNewFieldIsDisabled = errors.New("field is disabled")
	// ErrNewPowerOutagePeriodsMoreThanMaximum godoc
	ErrNewPowerOutagePeriodsMoreThanMaximum = errors.New("power outage periods are more than maximum")
	// ErrNewPowerOutagePeriodInvalid godoc
	ErrNewPowerOutagePeriodInvalid = errors.New("power outage period is invalid")
	// ErrNewPowerOutagePeriodOngoing godoc
	ErrNewPowerOutagePeriodOngoing = errors.New("power outage period is ongoing")
	// ErrNewGatewayIDIsInvalid godoc
	ErrNewGatewayIDIsInvalid = errors.New("gateway id is invalid")
	// ErrNewGatewayIDIsUsed godoc
	ErrNewGatewayIDIsUsed = errors.New("gateway id is used")
)

// ErrNewKeyNotExist godoc
func ErrNewKeyNotExist(key string) error {
	return fmt.Errorf("Key %s does not exist", key)
}

// ErrNewKeyUnexpectedValue godoc
func ErrNewKeyUnexpectedValue(key string) error {
	return fmt.Errorf("Key %s value is unexpected", key)
}

// ErrNewUserExpiration godoc
func ErrNewUserExpiration(expirationDate time.Time) error {
	return fmt.Errorf("User is expired on %s, please contact admin to extend", expirationDate)
}

// ErrNewUserLocked godoc
func ErrNewUserLocked(passwordLockCount int) error {
	return fmt.Errorf("PasswordRetryCount: over %d tries, please contact admin to unlock", passwordLockCount)
}
