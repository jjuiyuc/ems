package e

// Response code
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrAuthNoHeader      = 20000
	ErrAuthInvalidHeader = 20001
	ErrAuthTokenParse    = 20002
	ErrAuthTokenGen      = 20003
	ErrAuthLogin         = 20004

	ErrToken = 40000
)
