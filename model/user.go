package model

import (
  uuid "github.com/google/uuid"
)

type User struct {
  Id uuid.UUID
  Username string
  Password string
  Boxes []Box
}

func InitUser(username, password string) *User {
  return &User{
    Id: uuid.New(),
    Username: username,
    Password: password,
  }
}
