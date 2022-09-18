package model

type Session struct {
	UUID             string `gorm:"primary_key" json:"uuid,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	AccessTokenUUID  string `json:"access_token_uuid,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	RefreshTokenUUID string `json:"refresh_token_uuid,omitempty"`
}

func (Session) TableName() string { return "sessions" }
