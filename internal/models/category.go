package models

// Category model of ads.
type Category struct {
	BaseModel
	Title string `gorm:"unique"`
	AdID  uint
}
