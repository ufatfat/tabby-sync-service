package models

type OAuthBinding struct {
	BindingID       uint64 `gorm:"->:false;primaryKey;"`
	UserID          uint64
	OAuthType       uint8  `gorm:"column:oauth_type"`
	OAuthTypeString string `gorm:"column:oauth_type_string"`
	OAuthID         string `gorm:"column:oauth_id"`
}
