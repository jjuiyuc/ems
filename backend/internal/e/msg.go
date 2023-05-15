package e

var msgFlags = map[int]string{
	Success:                             "ok",
	Error:                               "fail",
	InvalidParams:                       "invalid parameters",
	ErrAccountGroupNameOnSameLevelExist: "account group name exists on the same level",
}

// GetMsg returns the information based on code
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[Error]
}
