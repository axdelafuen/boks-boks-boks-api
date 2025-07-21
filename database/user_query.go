package database

import (
	"main/model"

	"gorm.io/gorm"
)

func SelectUser(db *gorm.DB, username, password string) ([]model.User, error) {
	var user []model.User
	result := db.Where("username = ?", username).Where("password = ?", password).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func InsertUser(db *gorm.DB, newUser *model.User) error {
	return db.Create(&newUser).Error
}

func CheckUserOwnLabel(db *gorm.DB, userId, labelId string) ([]string, error) {
	var labelsID []string

	if err := db.Table("users_labels").Where("userid = ?", userId).Where("labelid = ?", labelId).Select("labelid").Find(&labelsID).Error; err != nil {
		return nil, err
	}

	return labelsID, nil
}
