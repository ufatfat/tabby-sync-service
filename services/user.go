package services

import "tabby-sync/models"

func GetUser(auth string) (user models.GetUser, err error) {
	err = db.Table("users").Select("user_id, username as name").Where("password = ?", auth).Scan(&user).Error
	return
}
