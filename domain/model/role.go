package model

type Role struct {
	ID   uint   `gorm:"primary_key" json:"id,omitempty"`
	Role string `json:"role,omitempty"`
}

func (Role) TableName() string { return "roles" }

const (
	RoleUserID = 1
	RoleAdvertisementID = 2
	RoleCommentID = 3
	RoleUserToRoleID = 4
	RoleAdminID = 5
)
