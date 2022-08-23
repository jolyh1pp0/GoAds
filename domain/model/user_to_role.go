package model

type UserRole struct {
	ID     uint   `gorm:"primary_key" json:"id,omitempty"`
	User   User   `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Role   Role   `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
	RoleID int    `json:"role_id,omitempty"`
}

func (UserRole) TableName() string { return "user_to_roles" }

type UserRoleResponseData struct {
	ID     uint   `gorm:"primary_key" json:"id,omitempty"`
	User   GetUsersResponseData `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	UserID string `json:"-"`
	Role   Role   `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
	RoleID int    `json:"-"`
}

func (UserRoleResponseData) TableName() string { return "user_to_roles"}
