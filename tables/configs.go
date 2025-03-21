package tables

type Config struct {
	ConfigID            uint64 `gorm:"primaryKey;autoIncrement"`
	Name                string `gorm:"size:255;notNull;default:'';"`
	Content             string `gorm:"size:8192;notNull;default:'';"`
	UserID              uint64 `gorm:"index:idx_user;notNull;default:0;"`
	LastUsedWithVersion string `gorm:"size:255;notNull;default:'';"`
	Deleted             uint8  `gorm:"notNull;default:0;index:idx_deleted;"`
	CreatedTimestamp    int64  `gorm:"type:timestamp default current_timestamp not null;"`
	UpdatedTimestamp    int64  `gorm:"type:timestamp default current_timestamp not null on update current_timestamp;"`
}
