package models

const (
	PublishedStatus = 1
	RejectedStatus  = 2
	PendingStatus   = 3
)

// Ad is the base ad model of the service.
type Ad struct {
	BaseModel
	Title       string `gorm:"unique"`
	Description string
	Status      int
	Image       string
	UserID      uint
	User        User
	Categories  []Category
}
