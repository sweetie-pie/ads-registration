package models

// Category model of ads.
type Category struct {
	BaseModel
	Title string
	AdID  uint
}
