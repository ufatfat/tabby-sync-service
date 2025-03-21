package services

import (
	"tabby-sync/models"
	"tabby-sync/utils"
)

func GetUser(auth string) (user models.GetUser, err error) {
	err = db.Table("users").Select("user_id, username as name").Where("token = ?", auth).Scan(&user).Error
	return
}

func NewUser(oauthType uint8, oauthTypeString, username, email, token string, oauthID any) (id uint64, err error) {
	tx := db.Begin()
	if token == "" {
		token = utils.GenRandString(32, 32)
	}
	userData := models.CreateUser{
		Username: username,
		Email:    email,
		Token:    token,
	}
	if err = tx.Table("users").Create(&userData).Error; err != nil {
		tx.Rollback()
		return
	}
	if oauthType == 0 {
		tx.Commit()
		return userData.UserID, nil
	}
	oauthIDString, err := utils.ToString(oauthID)
	if err != nil {
		tx.Rollback()
		return
	}
	bindingData := models.OAuthBinding{
		UserID:          userData.UserID,
		OAuthType:       oauthType,
		OAuthTypeString: oauthTypeString,
		OAuthID:         oauthIDString,
	}
	if err = tx.Table("oauth_bindings").Create(&bindingData).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return userData.UserID, nil
}
