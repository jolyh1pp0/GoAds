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
	VerifiedType string     `json:"verified_type,omitempty"`
	GoogleSecret string     `json:"google_secret,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func (User) TableName() string { return "users" }

type UserLogin struct {
	Email    string `json:"email,omitempty" mod:"trim" validate:"required"`
	Password string `json:"password,omitempty" mod:"trim" validate:"required"`
}

func (UserLogin) TableName() string { return "users" }

type UserRegister struct {
	ID           string `gorm:"primary_key" json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty" mod:"trim" validate:"required"`
	LastName     string `json:"last_name,omitempty" mod:"trim" validate:"required"`
	Email        string `json:"email,omitempty" mod:"trim" validate:"required"`
	Password     string `json:"password,omitempty" mod:"trim" validate:"required"`
	Phone        string `json:"phone,omitempty" mod:"trim" validate:"required"`
	VerifiedType string `json:"verified_type,omitempty" mod:"trim" validate:"required"`
	TimeZone     string `json:"time_zone,omitempty" mod:"trim" validate:"required"`
}

func (UserRegister ) TableName() string { return "users" }

type GetUsersResponseData struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func (GetUsersResponseData) TableName() string { return "users" }
