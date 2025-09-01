package service

import (
	"fmt"

	"main/database"
	"main/dto"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetUser(userId string) (*dto.UserResponse, error) {
	user, err := database.SelectUserById(s.db, userId)
	if err != nil {
		return nil, fmt.Errorf("Could not fetch user: %w", err)
	}

	return &dto.UserResponse{Id: user.Id.String(), Username: user.Username}, nil
}

func (s *UserService) GetUserWithUsername(userId, username string) (*dto.UserResponse, error) {
	users, err := database.SelectUserWithIdAndUsername(s.db, userId, username)
	if err != nil {
		return nil, fmt.Errorf("Could not fecth user: %w", err)
	}

	if len(*users) == 0 {
		return nil, fmt.Errorf("No user with username %s accessible with your token", username)
	}

	return &dto.UserResponse{Username: (*users)[0].Username}, nil
}

func (s *UserService) GetUserMetadata(username string) (*dto.UserMetadataResponse, error) {
	user, err := database.SelectUser(s.db, username)

	if err != nil {
		return nil, fmt.Errorf("Could not fecth user: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("No user with username %s accessible with your token", username)
	}

	boxCount, err := database.CountUserBoxes(s.db, user.Id.String())
	if err != nil {
		return nil, fmt.Errorf("Could not count user boxes: %w", err)
	}

	itemCount, err := database.CountUserItems(s.db, user.Id.String())
	if err != nil {
		return nil, fmt.Errorf("Could not count user items: %w", err)
	}

	labelCount, err := database.CountUserLabels(s.db, user.Id.String())
	if err != nil {
		return nil, fmt.Errorf("Could not count user labels: %w", err)
	}

	return &dto.UserMetadataResponse{
		UserResponse: dto.UserResponse{
			Id:       user.Id.String(),
			Username: user.Username,
		},
		TotalBoxes:  int(boxCount),
		TotalItems:  int(itemCount),
		TotalLabels: int(labelCount),
	}, nil
}
