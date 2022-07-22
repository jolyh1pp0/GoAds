package model

import "time"

type Comment struct {
	ID              string     `gorm:"primary_key" json:"id,omitempty"`
	AdvertisementId string     `json:"advertisement_id,omitempty"`
	Content         string     `json:"content,omitempty"`
	UserID          string     `json:"user_id,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

func (Comment) TableName() string { return "comments" }
