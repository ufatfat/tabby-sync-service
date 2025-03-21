package models

type GetUser struct {
	UserID uint64 `json:"-"`
	Name   string `json:"name"`
}

type CreateUser struct {
	UserID   uint64 `gorm:"->:false;primaryKey;"`
	Username string
	Email    string
	Token    string
}
