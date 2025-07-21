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

func (s *LabelService) GetLabels(userId string) (*[]dto.LabelResponse, error) {
	var res []dto.LabelResponse
	labels, err := database.SelectLabels(s.db, userId)
	if err != nil {
		return nil, fmt.Errorf("can't fetch data from db: %w", err)
	}

	for _, label := range labels {
		res = append(res, dto.LabelResponse{
			Id:    label.Id.String(),
			Title: label.Title,
			Color: label.Color,
		})
	}

	return &res, nil
}

func (s *LabelService) AddLabelToItem(userId, boxId, itemId, labelId string) error {
	boxesID, err := database.CheckBoxExist(s.db, userId, boxId)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	if len(boxesID) == 0 {
		return fmt.Errorf("box %s is not related to user %s", userId, boxId)
	}

	labelsID, err := database.CheckUserOwnLabel(s.db, userId, labelId)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	if len(labelsID) == 0 {
		return fmt.Errorf("user %s can't access label %s", userId, labelId)
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.InsertLinkItemLabel(tx, itemId, labelId); err != nil {
		tx.Rollback()
		return fmt.Errorf("error while inserting: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
