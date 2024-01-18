package models

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Status      int
	Image       string
	UserID      uint
	AdminID     uint
}
