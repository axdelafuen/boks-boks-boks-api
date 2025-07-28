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

func DeleteUserBoxLink(db *gorm.DB, userID, boxID string) error {
	return db.Table("users_boxes").Where("userid = ?", userID).Where("boxid= ?", boxID).Delete(nil).Error
}

func DeleteBoxItemLink(db *gorm.DB, boxID, itemID string) error {
	return db.Table("boxes_items").Where("boxid = ?", boxID).Where("itemid = ?", itemID).Delete(nil).Error
}

func InsertLinkUserLabel(db *gorm.DB, userId, labelId string) error {
	return db.Table("users_labels").Create(map[string]interface{}{
		"userid":  userId,
		"labelid": labelId,
	}).Error
}

func InsertLinkItemLabel(db *gorm.DB, itemId, labelId string) error {
	return db.Table("items_labels").Create(map[string]interface{}{
		"itemid":  itemId,
		"labelid": labelId,
	}).Error
}

func DeleteItemLabelsLink(db *gorm.DB, itemId string) error {
	return db.Table("items_labels").Where("itemid = ?", itemId).Delete(nil).Error
}

func DeleteItemLabelLinks(db *gorm.DB, itemId string, labelsIds []string) error {
	return db.Table("items_labels").Where("itemid = ?", itemId).Where("labelid IN ?", labelsIds).Delete(nil).Error
}

func InsertItemLabelLinks(db *gorm.DB, itemId string, labelsIds []string) error {
	for _, id := range labelsIds {
		if err := InsertLinkItemLabel(db, itemId, id); err != nil {
			return err
		}
	}

	return nil
}
