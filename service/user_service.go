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

func (s *UserService) GetUser(userId, username string) (*dto.UserResponse, error) {
	users, err := database.SelectUserWithId(s.db, userId, username)
	if err != nil {
		return nil, fmt.Errorf("Could not fecth user: %w", err)
	}

	if len(*users) == 0 {
		return nil, fmt.Errorf("No user with username %s accessible with your token", username)
	}

	return &dto.UserResponse{Username: (*users)[0].Username}, nil
}
