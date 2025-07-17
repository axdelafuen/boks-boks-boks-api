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
