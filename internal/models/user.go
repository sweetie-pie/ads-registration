package models

const (
	AccessLevelViewer = 1
	AccessLevelWriter = 2
	AccessLevelAdmin  = 3
)

type User struct {
	BaseModel
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	AccessLevel int
}
