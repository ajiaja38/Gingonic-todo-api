package config

import (
	"fmt"
	"learning-gin/src/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDatabase() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Todo{})

	return db, nil
}
