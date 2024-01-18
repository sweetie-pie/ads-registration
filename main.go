package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asaldelkhosh/ads-registration/internal/http"
	"github.com/asaldelkhosh/ads-registration/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func readDatabaseCredentialsFromEnv() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func migrateDatabaseModels(db *gorm.DB) error {
	classes := []interface{}{
		&models.User{},
		&models.Admin{},
		&models.Ad{},
		&models.Category{},
	}

	for _, class := range classes {
		if err := db.AutoMigrate(class); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	// connect to mysql database
	db, err := gorm.Open(mysql.Open(readDatabaseCredentialsFromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	// migrate models
	if err := migrateDatabaseModels(db); err != nil {
		log.Fatal("error migrating models", err)
	}

	// create image directory
	if _, err := os.Stat("/images"); os.IsNotExist(err) {
		if er := os.Mkdir("images", 0750); er != nil {
			log.Fatal("error creating images dir", er)
		}
	}

	// create new http handler
	httpHandler := http.HTTP{
		DB: db,
	}

	// register and start http server
	if err := httpHandler.Register(os.Getenv("HTTP_PORT")); err != nil {
		log.Fatal("error registering the http server", err)
	}
}
