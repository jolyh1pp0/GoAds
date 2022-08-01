package model

type User struct {
	ID           string     `gorm:"primary_key" json:"id,omitempty"`
	Email        string     `json:"email,omitempty"`
	Password     string     `json:"password,omitempty"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	TimeZone     string     `json:"-"`
	Phone        string     `json:"phone,omitempty"`
	Disabled     bool       `json:"-"`
	VerifiedType string     `json:"verified_type"`
	GoogleSecret string     `json:"-"`
}

func (User) TableName() string { return "users" }

