package model

import "time"

type Comment struct {
	ID              uint          `gorm:"primary_key" json:"id,omitempty"`
	AdvertisementId uint          `json:"advertisement_id,omitempty"`
	Content         string        `json:"content,omitempty"`
	User 	        User	      `json:"user,omitempty"`
	UserID          string        `json:"user_id"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
}

func (Comment) TableName() string { return "comments" }

