package service

import (
	"fmt"

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
		res = append(res, dto.ItemResponse{
			Id:     item.Id.String(),
			Title:  item.Title,
			Amount: item.Amount,
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

	newItem := model.InitItem(req.Title, req.Amount)

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

	if err := database.InsertLinkBoxItem(tx, boxIdDb[0], newItem.Id.String()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to link item to box: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Convert to DTO format before returning
	itemResponse := &dto.ItemResponse{
		Id:     newItem.Id.String(),
		Title:  newItem.Title,
		Amount: newItem.Amount,
	}

	return itemResponse, nil
}
