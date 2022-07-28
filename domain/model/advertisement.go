package model

import "time"

type Advertisement struct {
	ID          uint       `gorm:"primary_key" json:"id,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       int        `json:"price,omitempty"`
	Photo_1     string     `json:"photo_1,omitempty"`
	Photo_2     string     `json:"photo_2,omitempty"`
	Photo_3     string     `json:"photo_3,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (Advertisement) TableName() string { return "advertisements" }
