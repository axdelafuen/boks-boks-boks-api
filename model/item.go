package model

import (
	uuid "github.com/google/uuid"
)

type Item struct {
	Id     uuid.UUID
	Title  string
	Amount int
	Labels []Label `gorm:"-"`
}

func InitItem(title string, amount int) *Item {
	var i Item
	i.Id = uuid.New()
	i.Title = title
	i.Amount = amount

	return &i
}

func InitItemWithLabels(title string, amount int, labels []Label) *Item {
	var i Item
	i.Id = uuid.New()
	i.Title = title
	i.Amount = amount
	i.Labels = labels

	return &i
}
