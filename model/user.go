package model

import (
	"main/utils"

	uuid "github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;not null"`
	Username string    `json:"username" gorm:"not null"`
	Password string    `gorm:"not null"`
	Boxes    []Box     `gorm:"-"`
}

func InitUser(username, password string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       uuid.New(),
		Username: username,
		Password: hashedPassword,
	}, nil
}
