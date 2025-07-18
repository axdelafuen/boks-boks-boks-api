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

func (s *BoxService) DeleteBox(userID uuid.UUID, boxID string) error {
	boxesID, err := database.CheckBoxExist(s.db, userID.String(), boxID)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	if len(boxesID) == 0 {
		return fmt.Errorf("box %s is not related to user %s", userID.String(), boxID)
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := database.DeleteUserBoxLink(tx, userID.String(), boxID); err != nil {
		return fmt.Errorf("error while deleting user<->box link: %w", err)
	}

	items, err := database.SelectItems(tx, boxID)
	if err != nil {
		return fmt.Errorf("can't fecth datas: %w", err)
	}

	if len(*items) != 0 {
		for _, item := range *items {
			if err := database.DeleteBoxItemLink(tx, boxID, item.Id.String()); err != nil {
				return fmt.Errorf("error while deleting box<->item link: %w", err)
			}
		}
	}

	if err := database.DeleteBox(tx, boxID); err != nil {
		return fmt.Errorf("error while deleting box: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
