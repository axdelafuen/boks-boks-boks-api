package model

import (
	uuid "github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey;not null"`
	Username string    `json:"username" gorm:"not null"`
	Password string    `gorm:"not null"`
	Boxes    []Box     `gorm:"-"`
}

func InitUser(username, password string) *User {
	return &User{
		Id:       uuid.New(),
		Username: username,
		Password: password,
	}
}
