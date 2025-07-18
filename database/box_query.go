package database

import (
	"main/model"

	"gorm.io/gorm"
)

func SelectBoxes(db *gorm.DB, userID string) ([]model.Box, error) {
	var boxes []model.Box
	var boxesID []string

	if err := db.Table("users_boxes").Where("userid = ?", userID).Select("boxid").Find(&boxesID).Error; err != nil {
		return nil, err
	}

	if len(boxesID) == 0 {
		return boxes, nil
	}

	if err := db.Where("id IN ?", boxesID).Find(&boxes).Error; err != nil {
		return nil, err
	}

	return boxes, nil
}

func InsertBox(db *gorm.DB, box *model.Box) error {
	return db.Create(box).Error
}

func CheckBoxExist(db *gorm.DB, userID, boxID string) ([]string, error) {
	var boxIds []string

	if err := db.Table("users_boxes").Where("userid = ?", userID).Where("boxid = ?", boxID).Select("boxid").Find(&boxIds).Error; err != nil {
		return nil, err
	}

	return boxIds, nil
}

func DeleteBox(db *gorm.DB, id string) error {
	return db.Where("id = ?", id).Delete(&model.Box{}).Error
}

func UpdateBox(db *gorm.DB, id, title string) error {
	return db.Model(&model.Box{}).Where("id = ?", id).UpdateColumn("title", title).Error
}
