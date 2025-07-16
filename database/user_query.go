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
