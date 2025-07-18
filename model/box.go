package model

import (
	uuid "github.com/google/uuid"
)

type Box struct {
	Id    uuid.UUID `gorm:"primaryKey; not null" json:"id"`
	Title string    `gorm:"not null" json:"title"`
	Items []Item    `gorm:"-" json:"items"`
}

func InitBox(title string) *Box {
	var b Box
	b.Id = uuid.New()
	b.Title = title

	return &b
}
