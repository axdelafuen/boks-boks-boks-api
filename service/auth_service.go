package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"main/database"
	"main/dto"
	"main/middleware"
	"main/model"
	"main/utils"
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
	user, err := database.SelectUser(s.db, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := middleware.GenerateJWT(user.Id, s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.AuthResponse{
		Token: token,
	}, nil
}

func (s *AuthService) Register(req *dto.AuthRequest) error {
	// Validate password strength
	if !utils.IsValidPasswordLength(req.Password) {
		return errors.New("password must be at least 8 characters long")
	}

	newUser, err := model.InitUser(req.Username, req.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := database.InsertUser(s.db, newUser); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
