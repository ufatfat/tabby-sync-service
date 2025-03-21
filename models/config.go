package models

import "time"

type GetConfigList []GetConfig

type GetConfig struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type CreateConfig struct {
	ID      uint64 `json:"id,omitempty" gorm:"primaryKey;column:config_id;"`
	UserID  uint64 `json:"user_id" gorm:"unique;column:user_id;"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type ConfigResult struct {
	ID                  uint64    `json:"id,omitempty" gorm:"primaryKey;column:config_id;"`
	Name                string    `json:"name"`
	Content             string    `json:"content"`
	CreatedAt           time.Time `json:"created_at"`
	ModifiedAt          time.Time `json:"modified_at"`
	LastUsedWithVersion string    `json:"last_used_with_version"`
	User                uint64    `json:"user" gorm:"column:user_id"`
}

type UpdateConfig struct {
	LastUsedWithVersion string `json:"last_used_with_version"`
	Content             string `json:"content,omitempty"`
}
