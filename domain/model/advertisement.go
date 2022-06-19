package model

import "time"

type Advertisement struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float32    `json:"price"`
	Photo_1     string     `json:"photo_1"`
	Photo_2     string     `json:"photo_2"`
	Photo_3     string     `json:"photo_3"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (Advertisement) TableName() string { return "advertisements" }
