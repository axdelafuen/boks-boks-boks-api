package service

import (
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/model"
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

func (s *BoxService) CreateBox(userID uuid.UUID, req *dto.CreateBoxRequest) (*model.Box, error) {
	newBox := model.InitBox(req.Title)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertBox(tx, newBox); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create box: %w", err)
	}

	if err := database.InsertLinkUserBox(tx, userID.String(), newBox.Id.String()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to link box to user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newBox, nil
}
