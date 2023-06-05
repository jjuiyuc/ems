package app

// EnableFieldBody godoc
type EnableFieldBody struct {
	Enable *bool `form:"enable" binding:"required"`
}

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

// CreateUserBody godoc
type CreateUserBody struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required,max=50"`
	Name     string `form:"name" binding:"required,max=20"`
	GroupID  int    `form:"groupID" binding:"required"`
}

// UpdateUserBody godoc
type UpdateUserBody struct {
	Password string `form:"password" binding:"max=50"`
	Name     string `form:"name" binding:"max=20"`
	GroupID  int    `form:"groupID"`
	Unlock   bool   `form:"unlock"`
}
