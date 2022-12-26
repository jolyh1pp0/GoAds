package model

import "time"

type Advertisement struct {
	ID          uint       `gorm:"primary_key" json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       int        `json:"price,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	User        User       `json:"user,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
}

func (Advertisement) TableName() string { return "advertisements" }

type AdvertisementsRequestData struct {
	Title       string               `json:"title,omitempty" validate:"required,max=200"`
	Description string               `json:"description,omitempty" validate:"required,max=1000"`
	Price       int                  `json:"price,omitempty" validate:"required"`
	UserID      string               `json:"-"`
	User        GetUsersResponseData `gorm:"foreignKey:UserID;references:ID" json:"author,omitempty"`
}

func (AdvertisementsRequestData) TableName() string { return "advertisements" }

type AdvertisementsUpdateRequestData struct {
	Title       string `json:"title,omitempty" validate:"max=200"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	Price       int    `json:"price,omitempty"`
}

func (AdvertisementsUpdateRequestData) TableName() string { return "advertisements" }

type GetAdvertisementsResponseData struct {
	ID          uint                      `gorm:"primary_key" json:"id,omitempty"`
	Title       string                    `json:"title,omitempty"`
	Description string                    `json:"description,omitempty"`
	Price       int                       `json:"price,omitempty"`
	Gallery     []GetGalleryResponseData  `gorm:"foreignKey:AdvertisementId" json:"gallery,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
	UpdatedAt   *time.Time                `json:"updated_at,omitempty"`
	UserID      string                    `json:"-"`
	User        GetUsersResponseData      `gorm:"foreignKey:UserID;references:ID" json:"author,omitempty"`
	Comments    []GetCommentsResponseData `gorm:"foreignKey:AdvertisementId" json:"comments,omitempty"`
}

func (GetAdvertisementsResponseData) TableName() string { return "advertisements" }
