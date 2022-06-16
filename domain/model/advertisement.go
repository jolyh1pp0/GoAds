package model

import "time"

type Advertisement struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Title       string     `json:"name"`
	Description string     `json:"age"`
	Price       float32    `json:"price"`
	Photo_1     string     `json:"photo___1"`
	Photo_2     string     `json:"photo___2"`
	Photo_3     string     `json:"photo___3"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (Advertisement) TableName() string { return "advertisements" }
