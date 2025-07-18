package service

import (
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
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
			Id:    item.Id.String(),
			Title: item.Title,
		})
	}

	return &res, nil
}
