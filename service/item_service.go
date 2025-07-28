package service

import (
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/model"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{
		db: db,
	}
}

func (s *ItemService) GetItems(userID string, boxID string) (*[]dto.ItemResponse, error) {
	var res []dto.ItemResponse

	boxIdDb, err := database.CheckBoxExist(s.db, userID, boxID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch box: %w", err.Error())
	}

	if len(boxIdDb) == 0 {
		return nil, fmt.Errorf("Box id no found in db: %w", err.Error())
	}

	items, err := database.SelectItems(s.db, boxIdDb[0])
	if err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err.Error())
	}

	for _, item := range *items {
		// get labels for each item
		labels, err := database.SelectItemsLabels(s.db, item.Id.String())
		if err != nil {
			return nil, fmt.Errorf("failed to fetch labels: %w", err.Error())
		}

		res = append(res, dto.ItemResponse{
			Id:     item.Id.String(),
			Title:  item.Title,
			Amount: item.Amount,
			Labels: *labels,
		})
	}

	return &res, nil
}

func (s *ItemService) CreateItem(userID string, boxID string, req *dto.CreateItemRequest) (*dto.ItemResponse, error) {
	boxIdDb, err := database.CheckBoxExist(s.db, userID, boxID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify box ownership: %w", err)
	}

	if len(boxIdDb) == 0 {
		return nil, fmt.Errorf("box not found or access denied")
	}

	var labels []model.Label
	for _, label := range req.Labels {
		id, err := uuid.Parse(label.Id)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing label id: %w", err)
		}
		labels = append(labels, model.Label{
			Id:          id,
			Title:       label.Title,
			Description: label.Description,
			Color:       label.Color,
		})
	}
	newItem := model.InitItemWithLabels(req.Title, req.Amount, labels)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertItem(tx, newItem); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	for _, label := range req.Labels {
		if err := database.InsertLinkItemLabel(tx, newItem.Id.String(), label.Id); err != nil {
			return nil, fmt.Errorf("failed to link label to label: %w", err)
		}
	}

	if err := database.InsertLinkBoxItem(tx, boxIdDb[0], newItem.Id.String()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to link item to box: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	var labelsDto []dto.LabelResponse
	for _, l := range newItem.Labels {
		labelsDto = append(labelsDto, dto.LabelResponse{
			Id:          l.Id.String(),
			Title:       l.Title,
			Description: l.Description,
			Color:       l.Color,
		})
	}

	itemResponse := &dto.ItemResponse{
		Id:     newItem.Id.String(),
		Title:  newItem.Title,
		Amount: newItem.Amount,
		Labels: labelsDto,
	}

	return itemResponse, nil
}

func (s *ItemService) DeleteItem(userId, boxId, itemId string) error {
	boxIdDb, err := database.CheckBoxExist(s.db, userId, boxId)
	if err != nil {
		return fmt.Errorf("failed to verify box ownership: %w", err)
	}

	if len(boxIdDb) == 0 {
		return fmt.Errorf("box not found or access denied")
	}

	itemIdDb, err := database.CheckBoxOwnItem(s.db, boxId, itemId)
	if err != nil {
		return fmt.Errorf("failed to verify item ownership: %w", err)
	}

	if len(itemIdDb) == 0 {
		return fmt.Errorf("item not found or access denied")
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.DeleteBoxItemLink(tx, boxId, itemId); err != nil {
		tx.Rollback()
		return fmt.Errorf("fail while deleting box<->item link: %w", err)
	}

	if err := database.DeleteItemLabelsLink(tx, itemId); err != nil {
		tx.Rollback()
		return fmt.Errorf("fail while deleting item<->label link: %w", err)
	}

	if err := database.DeleteItemWithId(tx, itemId); err != nil {
		tx.Rollback()
		return fmt.Errorf("fail while deleting item: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *ItemService) UpdateItem(userId, boxId string, req *dto.UpdateItemRequest) (*dto.ItemResponse, error) {
	boxIdDb, err := database.CheckBoxExist(s.db, userId, boxId)
	if err != nil {
		return nil, fmt.Errorf("failed to verify box ownership: %w", err)
	}

	if len(boxIdDb) == 0 {
		return nil, fmt.Errorf("box not found or access denied: %w", err)
	}

	itemIdDb, err := database.CheckBoxOwnItem(s.db, boxId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to verify item ownership: %w", err)
	}

	if len(itemIdDb) == 0 {
		return nil, fmt.Errorf("item not found or access denied")
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newItem, err := database.UpdateItem(tx, req.Id, req.Title, req.Amount)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error while updating item")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newItem, nil
}
