package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/middleware"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(req *dto.AuthRequest) (*dto.AuthResponse, error) {
	users, err := database.SelectUser(s.db, req.Username, req.Password)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	if len(users) == 0 {
		return nil, errors.New("invalid credentials")
	}

	token, err := middleware.GenerateJWT(users[0].Id, s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.AuthResponse{
		Token: token,
	}, nil
}
