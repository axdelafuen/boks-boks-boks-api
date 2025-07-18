package database

import (
	"gorm.io/gorm"
)

func InsertLinkUserBox(db *gorm.DB, userId string, boxId string) error {
	return db.Table("users_boxes").Create(map[string]interface{}{
		"userid": userId,
		"boxid":  boxId,
	}).Error
}

func InsertLinkBoxItem(db *gorm.DB, boxId string, itemId string) error {
	return db.Table("boxes_items").Create(map[string]interface{}{
		"boxid":  boxId,
		"itemid": itemId,
	}).Error
}
