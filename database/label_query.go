package database

import (
	"main/model"

	"gorm.io/gorm"
)

func InsertLabel(db *gorm.DB, label *model.Label) error {
	return db.Create(label).Error
}
