package main

import (
	"fmt"
	"log"
	"os"

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
}
