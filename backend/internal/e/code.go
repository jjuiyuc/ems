package e

// Response code
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrAuthNoHeader              = 20000
	ErrAuthInvalidHeader         = 20001
	ErrAuthTokenParse            = 20002
	ErrAuthTokenGen              = 20003
	ErrAuthUserNotExist          = 20004
	ErrAuthUserExpirated         = 20005
	ErrAuthUserLocked            = 20006
	ErrAuthPasswordNotMatch      = 20007
	ErrAuthPolicyLoad            = 20008
	ErrAuthPermissionCheck       = 20009
	ErrAuthPermissionNotAllow    = 20010
	ErrAuthFrontendPermissionGen = 20011

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

	ErrNameUpdate                       = 60000
	ErrPasswordUpdate                   = 60001
	ErrAccountGroupsGen                 = 60002
	ErrAccountGroupNameOnSameLevelExist = 60003
	ErrAccountGroupCreate               = 60004
	ErrAccountGroupGen                  = 60005
	ErrOwnAccountGroupModifiedNotAllow  = 60006
	ErrAccountGroupUpdate               = 60007
	ErrAccountGroupHasSubGroup          = 60008
	ErrAccountGroupHasUser              = 60009
	ErrAccountGroupDelete               = 60010
	ErrAccountUsersGen                  = 60011
	ErrAccountUsernameExist             = 60012
	ErrAccountUserCreate                = 60013
	ErrAccountUserUpdate                = 60014
	ErrOwnAccountDeletedNotAllow        = 60015
	ErrAccountUserDelete                = 60016
	ErrFieldsGen                        = 60017
	ErrDeviceModelsGen                  = 60018
	ErrFieldGen                         = 60019
	ErrFieldEnable                      = 60020
	ErrFieldGroupsUpdate                = 60021
	ErrFieldIsDisabled                  = 60022
	ErrDeviceSettingsSync               = 60023
	ErrBatterySettingsGen               = 60030
	ErrBatterySettingsUpdate            = 60031
	ErrMeterSettingsGen                 = 60032
	ErrMeterSettingsUpdate              = 60033
)
