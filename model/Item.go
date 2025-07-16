package model 

import (
  uuid "github.com/google/uuid"
)

type Item struct {
  Id uuid.UUID
  Title string
  Amount int
}

func InitItem(title string, amount int) {
 var i Item
 i.Id = uuid.New()
 i.Title = title
 i.Amount = amount

 return &i
}
