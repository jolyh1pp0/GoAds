package model

import "time"

type PasswordRecovery struct {
	ID        int        `gorm:"primary_key" json:"id,omitempty"`
	User      User       `json:"user,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	UserEmail string     `json:"email,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiresAt time.Time  `json:"expires_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (PasswordRecovery) TableName() string { return "passwords_recovery" }
