package database

import (
	"main/dto"
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

func SelectItemsLabels(db *gorm.DB, itemId string) (*[]dto.LabelResponse, error) {
	var labels []model.Label
	var labelsId []string
	var res []dto.LabelResponse

	if err := db.Table("items_labels").Where("itemid = ?", itemId).Select("labelid").Find(&labelsId).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id IN ?", labelsId).Find(&labels).Error; err != nil {
		return nil, err
	}

	for _, l := range labels {
		res = append(res, dto.LabelResponse{
			Id:    l.Id.String(),
			Title: l.Title,
			Color: l.Color,
		})
	}

	return &res, nil
}
