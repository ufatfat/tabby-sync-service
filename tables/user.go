package tables

type User struct {
	UserID                uint64 `gorm:"primaryKey;autoIncrement;"`
	Username              string `gorm:"size:255;notNull;default:'';index:idx_username,unique;index:idx_login"`
	Password              string `gorm:"size:255;notNull;default:'';index:idx_login;"`
	Email                 string `gorm:"size:255;notNull;default:'';index:idx_email;"`
	Token                 string `gorm:"size:64;notNull;default:'';index:idx_token,unique;"`
	LastSignedInTimestamp int64  `gorm:"type:timestamp default current_timestamp not null;"`
	LastSignedInIP        string `gorm:"size:255;notNull;default:0.0.0.0;"`
	CreatedTimestamp      int64  `gorm:"type:timestamp default current_timestamp not null;"`
	UpdatedTimestamp      int64  `gorm:"type:timestamp default current_timestamp not null on update current_timestamp;"`
}

type OAuthBinding struct {
	BindingID        uint64 `gorm:"primaryKey;autoIncrement;"`
	UserID           uint64 `gorm:"index:idx_user_oauth;index:idx_user_oauth_type,unique;"`
	OAuthType        uint8  `gorm:"index:idx_oauth,unique;index:idx_user_oauth_type,unique;"`
	OAuthTypeString  string `gorm:"size:255;index:idx_oauth_str,unique;"`
	OAuthID          string `gorm:"size:255;index:idx_oauth,unique;index:idx_oauth_str,unique;"`
	CreatedTimestamp int64  `gorm:"type:timestamp default current_timestamp not null;"`
	UpdatedTimestamp int64  `gorm:"type:timestamp default current_timestamp not null on update current_timestamp;"`
}

func (OAuthBinding) TableName() string {
	return "oauth_bindings"
}
