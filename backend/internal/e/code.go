package e

// Response code
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrAuthNoHeader         = 20000
	ErrAuthInvalidHeader    = 20001
	ErrAuthTokenParse       = 20002
	ErrAuthTokenGen         = 20003
	ErrAuthUserNotExist     = 20004
	ErrAuthUserExpirated    = 20005
	ErrAuthUserLocked       = 20006
	ErrAuthPasswordNotMatch = 20007
	ErrAuthPolicyLoad       = 20008
	ErrAuthPolicyCheck      = 20009
	ErrAuthPolicyNotAllow   = 20010

	ErrPasswordToken = 30000
	// 30001
	ErrPasswordLost = 30002

	ErrToken = 40000

	ErrUserProfileGen               = 50000
	ErrDashboardDataGen             = 50001
	ErrBatteryPowerStateGen         = 50002
	ErrBatteryChargeVoltageStateGen = 50003
	ErrSolarPowerStateGen           = 50004
	ErrGridPowerStateGen            = 50005
	ErrTimeOfUseInfoGen             = 50006

	ErrNameUpdate     = 60000
	ErrPasswordUpdate = 60001
)
