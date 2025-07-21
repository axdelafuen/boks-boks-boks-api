package model

import (
	uuid "github.com/google/uuid"
)

type Label struct {
	Id    uuid.UUID
	Title string
	Color string
}

func InitLabel(title, color string) *Label {
	return &Label{Id: uuid.New(), Title: title, Color: color}
}
