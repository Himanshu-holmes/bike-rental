package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
type Bike struct {
	Id            string                
	OwnerName     string                
	Type          string               
	Make          string                 
	Serial        string                 
	}

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

	bike := new(Bike)
	err = db.AutoMigrate(bike)
	return db, err
}

func AttachCollection(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}
}
