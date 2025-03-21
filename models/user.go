package models

type GetUser struct {
	UserID uint64 `json:"-"`
	Name   string `json:"name"`
}
