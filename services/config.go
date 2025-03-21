package services

import (
	"tabby-sync/models"
)

func GetConfigList(userID uint64) (configs models.GetConfigList, err error) {
	err = db.Table("configs").Select("config_id as id, name, created_timestamp as created_at, updated_timestamp as modified_at").Where("user_id = ? and deleted = 0", userID).Find(&configs).Error
	return
}

func GetConfig(userID, configID uint64) (config models.GetConfig, err error) {
	err = db.Table("configs").Select("config_id as id, name, content, created_timestamp as created_at, updated_timestamp as modified_at").Where("config_id = ? and user_id = ? and deleted = 0", userID, configID).Scan(&config).Error
	return
}

func CreateConfig(data models.CreateConfig) (rst models.ConfigResult, err error) {
	if err = db.Table("configs").Create(&data).Error; err != nil {
		return
	}
	err = db.Table("configs").Where("config_id = ? and deleted = 0", data.ID).Scan(&rst).Error
	return
}

func UpdateConfig(userID, configID uint64, data models.UpdateConfig) (rst models.ConfigResult, err error) {
	if err = db.Table("configs").Where("config_id = ? and user_id = ? and deleted = 0", configID, userID).Updates(&data).Error; err != nil {
		return
	}
	err = db.Table("configs").Where("config_id = ? and user_id = ? and deleted = 0", configID, userID).Scan(&data).Error
	return
}

func DeleteConfig(userID, configID uint64) (err error) {
	err = db.Table("configs").Where("config_id = ? and user_id = ?", configID, userID).Update("deleted", true).Error
	return
}
