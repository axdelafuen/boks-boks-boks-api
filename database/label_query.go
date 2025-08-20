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

func DeleteLabel(db *gorm.DB, labelId string) error {
	return db.Where("id = ?", labelId).Delete(&model.Label{}).Error
}

func UpdateLabel(db *gorm.DB, id, title, description, color string) error {
	if err := db.Model(&model.Label{}).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "description": description, "color": color}).Error; err != nil {
		return err
	}

	return nil
}

func SelectLabelWithId(db *gorm.DB, id string) (*model.Label, error) {
	var label model.Label

	if err := db.Where("id = ?", id).First(&label).Error; err != nil {
		return nil, err
	}

	return &label, nil
}
