package app

// CreateGroupBody godoc
type CreateGroupBody struct {
	Name string `form:"name" binding:"required,max=20"`
	// TypeID 3 is "Area maintainer" and 4 is "Field owner"
	TypeID   int `form:"typeID" binding:"required,oneof=3 4"`
	ParentID int `form:"parentID" binding:"required"`
}

// UpdateGroupBody godoc
type UpdateGroupBody struct {
	Name string `form:"name" binding:"required,max=20"`
}
