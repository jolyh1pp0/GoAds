package model

import "time"

type Advertisement struct {
	ID          uint       `gorm:"primary_key" json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       int        `json:"price,omitempty"`
	Photo_1     string     `json:"photo_1,omitempty"`
	Photo_2     string     `json:"photo_2,omitempty"`
	Photo_3     string     `json:"photo_3,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	User        User       `json:"user,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
}

func (Advertisement) TableName() string { return "advertisements" }

type GetAdvertisementsResponseData struct {
	ID          uint                      `gorm:"primary_key" json:"id,omitempty"`
	Title       string                    `json:"title,omitempty"`
	Description string                    `json:"description,omitempty"`
	Price       int                       `json:"price,omitempty"`
	Photo_1     string                    `json:"photo_1,omitempty"`
	Photo_2     string                    `json:"photo_2,omitempty"`
	Photo_3     string                    `json:"photo_3,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
	UpdatedAt   *time.Time                `json:"updated_at,omitempty"`
	UserID      string                    `json:"-"`
	User        GetUsersResponseData      `gorm:"foreignKey:UserID;references:ID" json:"author,omitempty"`
	Comments    []GetCommentsResponseData `gorm:"foreignKey:AdvertisementId" json:"comments,omitempty"`
}

func (GetAdvertisementsResponseData) TableName() string { return "advertisements" }
