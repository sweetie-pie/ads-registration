package models

const (
	AccessLevelViewer = 1
	AccessLevelWriter = 2
	AccessLevelAdmin  = 3
)

type Admin struct {
	BaseModel
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	Active      bool
	AccessLevel int
}
