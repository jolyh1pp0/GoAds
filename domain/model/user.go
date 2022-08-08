package model

import "time"

type User struct {
	ID           string     `gorm:"primary_key" json:"id,omitempty"`
	Email        string     `json:"email,omitempty"`
	Password     string     `json:"password,omitempty"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	TimeZone     string     `json:"time_zone,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	Disabled     bool       `json:"disabled,omitempty"`
	VerifiedType string     `json:"verified_type"`
	GoogleSecret string     `json:"google_secret,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func (User) TableName() string { return "users" }
