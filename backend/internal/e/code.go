package e

// Response code
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// ErrAuthNoHeader      = 20000
	// ErrAuthInvalidHeader = 20001
	// ErrAuthTokenParse    = 20002
	// ErrAuthTokenTimeout  = 20003
	ErrAuthTokenGen = 20004
	ErrAuthLogin    = 20005
)
