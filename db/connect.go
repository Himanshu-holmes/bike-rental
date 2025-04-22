package db

import (
	"log"
	"os"

	"github.com/himanshuholmes/bikerental/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ParseEnvVars(key string) string {
	return os.Getenv(key)
}

func ConnectDB() (*gorm.DB, error) {
	 DB_URL := ParseEnvVars("DB_URL")
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{
		   Logger: logger.Default.LogMode(logger.Info), // for SQL logging
    DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	bike := new(models.BikeModel)
	err = db.AutoMigrate(bike)
	rentee := new(models.RenteeModel)
	err = db.AutoMigrate(rentee)
	return db, err
}

func AttachCollection(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}
}
