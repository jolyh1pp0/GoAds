package model

import "time"

type Session struct {
	ID                    string     `gorm:"primary_key" json:"id,omitempty"`
	User                  User       `json:"user,omitempty"`
	UserID                string     `json:"user_id,omitempty"`
	AccessTokenUUID       string     `json:"access_token_uuid,omitempty"`
	RefreshTokenUUID      string     `json:"refresh_token_uuid,omitempty"`
	RefreshTokenExpiresAt time.Time  `json:"refresh_token_expires_at,omitempty"`
	ExpiresAt             time.Time  `json:"expires_at,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
}

func (Session) TableName() string { return "sessions" }
