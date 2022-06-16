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

	ErrPasswordToken = 30000
	// ErrPasswordReset = 30001
	ErrPasswordLost = 30002

	ErrToken = 40000
)
