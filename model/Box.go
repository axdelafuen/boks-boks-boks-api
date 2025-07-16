package model 

import (
  uuid "github.com/google/uuid"
)

type Box struct {
  Id uuid.UUID
  Title string
  Items []Item
}

func InitBox(title string) *Box {
  var b Box
  b.Id = uuid.New()
  b.Title = title

  return &b
}
