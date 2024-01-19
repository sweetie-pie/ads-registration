package models

type Category struct {
	BaseModel
	Title string `gorm:"unique"`
	AdID  uint
}
