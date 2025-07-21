package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertLabel(db *gorm.DB, label *model.Label) error {
	return db.Create(label).Error
}

func SelectLabels(db *gorm.DB, userId string) ([]model.Label, error) {
	var labels []model.Label
	var labelsId []string

	if err := db.Table("users_labels").Where("userid = ?", userId).Select("labelid").Find(&labelsId).Error; err != nil {
		return nil, err
	}

	if len(labelsId) == 0 {
		return labels, nil
	}

	if err := db.Where("id IN ?", labelsId).Find(&labels).Error; err != nil {
		return nil, err
	}

	return labels, nil
}
