package app

// GroupURI godoc
type GroupURI struct {
	GroupID int64 `uri:"groupid" binding:"required"`
}
