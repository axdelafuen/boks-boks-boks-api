package service

import (
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/model"
)

type LabelService struct {
	db *gorm.DB
}

func NewLabelService(db *gorm.DB) *LabelService {
	return &LabelService{
		db: db,
	}
}

func (s *LabelService) CreateLabel(userId string, req *dto.CreateLabelRequest) (*model.Label, error) {
	newLabel := model.InitLabel(req.Title, req.Color)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertLabel(tx, newLabel); err != nil {
		return nil, fmt.Errorf("failed to create Label: %w", err)
	}

	if err := database.InsertLinkUserLabel(tx, userId, newLabel.Id.String()); err != nil {
		return nil, fmt.Errorf("failed to link user to label: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newLabel, nil
}
