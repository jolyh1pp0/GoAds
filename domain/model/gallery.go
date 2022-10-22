package model

import "time"

type Gallery struct {
	ID              uint          `gorm:"primary_key" json:"id,omitempty"`
	FileName        string        `json:"file_name,omitempty"`
	FilePath 		string 	      `json:"file_path,omitempty"`
	Advertisement   Advertisement `json:"advertisement,omitempty"`
	AdvertisementId uint          `json:"advertisement_id,omitempty"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
}

func (Gallery) TableName() string { return "gallery" }

type GetGalleryResponseData struct {
	ID              uint          `json:"-"`
	FilePath 		string 	      `json:"file_path"`
	Advertisement   Advertisement `json:"-"`
	AdvertisementId uint          `json:"-"`
}

func (GetGalleryResponseData) TableName() string { return "gallery" }

