package database

import (
	"main/model"

	"gorm.io/gorm"
)

func SelectBoxes(db *gorm.DB, userID string) ([]model.Box, error) {
	var boxes []model.Box
	var boxesID []string

	if err := db.Table("users_boxes").Where("userid = ?", userID).Select("tableid").Find(&boxesID).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id IN ?", boxesID).Find(&boxes).Error; err != nil {
		return nil, err
	}

	return boxes, nil
}

func InsertBox(db *gorm.DB, box *model.Box) error {
	return db.Create(box).Error
}
