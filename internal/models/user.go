package models

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Banned   bool
}
