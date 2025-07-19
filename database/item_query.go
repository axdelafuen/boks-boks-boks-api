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

func InsertItem(db *gorm.DB, item *model.Item) error {
	return db.Create(item).Error
}

func DeleteItem(db *gorm.DB, item *model.Item) error {
	return db.Where("id = ?", item.Id.String()).Delete(&model.Item{}).Error
}

func DeleteItems(db *gorm.DB, items *[]model.Item) error {
	for _, item := range *items {
		if err := DeleteItem(db, &item); err != nil {
			return err
		}
	}

	return nil
}
