package database

import (
	"main/model"

	"gorm.io/gorm"
)

func SelectItems(db *gorm.DB, boxID string) (*[]model.Item, error) {
	var itemIds []string
	var items []model.Item

	if err := db.Table("boxes_items").Where("boxid = ?", boxID).Select("itemid").Find(&itemIds).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id IN ?", itemIds).Find(&items).Error; err != nil {
		return nil, err
	}

	return &items, nil
}
