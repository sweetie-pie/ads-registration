package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel of our database.
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
