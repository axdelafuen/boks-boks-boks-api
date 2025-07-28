package database

import (
	"main/dto"
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

func DeleteItemWithId(db *gorm.DB, itemId string) error {
	return db.Where("id = ?", itemId).Delete(&model.Item{}).Error
}

func UpdateItemLabelsWithNew(db *gorm.DB, itemId string, newLabels []dto.LabelResponse) error {
	labels, err := SelectItemsLabels(db, itemId)
	if err != nil {
		return err
	}

	currentLabelsIds := extractIds(*labels)
	newLabelsIds := extractIds(newLabels)

	removed, toAdd := getAddRemoveLabel(currentLabelsIds, newLabelsIds)

	if err := DeleteItemLabelLinks(db, itemId, removed); err != nil {
		return err
	}

	if err := InsertItemLabelLinks(db, itemId, toAdd); err != nil {
		return err
	}

	return nil
}

func getAddRemoveLabel(sourceId, newId []string) (removed []string, toAdd []string) {
	srcIds := make(map[string]bool)
	newIds := make(map[string]bool)

	for _, id := range sourceId {
		srcIds[id] = true
	}

	for _, id := range newId {
		newIds[id] = true
	}

	for id := range srcIds {
		if !newIds[id] {
			removed = append(removed, id)
		}
	}

	for id := range newIds {
		if !srcIds[id] {
			toAdd = append(toAdd, id)
		}
	}

	return
}

func extractIds(labels []dto.LabelResponse) []string {
	ids := make([]string, 0, len(labels))
	for _, label := range labels {
		ids = append(ids, label.Id)
	}
	return ids
}

func UpdateItem(db *gorm.DB, id, title string, amount int, labels []dto.LabelResponse) (*dto.ItemResponse, error) {
	if err := db.Model(&model.Item{}).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "amount": amount}).Error; err != nil {
		return nil, err
	}

	var updatedItem model.Item
	if err := db.Where("id = ?", id).First(&updatedItem).Error; err != nil {
		return nil, err
	}

	if err := UpdateItemLabelsWithNew(db, id, labels); err != nil {
		return nil, err
	}

	labelsDb, err := SelectItemsLabels(db, id)
	if err != nil {
		return nil, err
	}

	var labelsDto []dto.LabelResponse
	if labelsDb != nil {
		labelsDto = *labelsDb
	}

	itemResponse := &dto.ItemResponse{
		Id:     updatedItem.Id.String(),
		Title:  updatedItem.Title,
		Amount: updatedItem.Amount,
		Labels: labelsDto,
	}

	return itemResponse, nil
}
