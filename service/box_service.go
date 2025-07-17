package service

import (
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
)

type BoxService struct {
	db *gorm.DB
}

func NewBoxService(db *gorm.DB) *BoxService {
	return &BoxService{
		db: db,
	}
}

func (s *BoxService) GetBoxes(userID string) (*[]dto.BoxResponse, error) {
	var res []dto.BoxResponse
	boxes, err := database.SelectBoxes(s.db, userID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch data from db: %w", err)
	}

	for _, box := range boxes {
		res = append(res, dto.BoxResponse{
			Id:    box.Id.String(),
			Title: box.Title,
		})
	}

	return &res, nil
}
