package tables

type User struct {
	UserID                uint64 `gorm:"primaryKey;autoIncrement;"`
	Username              string `gorm:"size:255;notNull;default:'';"`
	Password              string `gorm:"size:255;notNull;default:'';index:idx_password,unique;"`
	LastSignedInTimestamp int64  `gorm:"type:timestamp default current_timestamp not null;"`
	LastSignedInIP        string `gorm:"size:255;notNull;default:0.0.0.0;"`
	CreatedTimestamp      int64  `gorm:"type:timestamp default current_timestamp not null;"`
	UpdatedTimestamp      int64  `gorm:"type:timestamp default current_timestamp not null on update current_timestamp;"`
}
