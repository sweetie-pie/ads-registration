package models

const (
	PublishedStatus = 1
	RejectedStatus  = 2
	PendingStatus   = 3
)

type Ad struct {
	BaseModel
	Title       string `gorm:"unique"`
	Description string
	Status      int
	Image       string
	UserID      uint
	AdminID     uint
	Categories  []Category
}
