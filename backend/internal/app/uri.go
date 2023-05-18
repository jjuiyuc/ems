package app

// GroupURI godoc
type GroupURI struct {
	GroupID int64 `uri:"groupid" binding:"required"`
}

// UserURI godoc
type UserURI struct {
	UserID int64 `uri:"userid" binding:"required"`
}
