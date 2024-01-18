package http

import "gorm.io/gorm"

type HTTP struct {
	DB *gorm.DB
}

func (h HTTP) Register() error {
	return nil
}
