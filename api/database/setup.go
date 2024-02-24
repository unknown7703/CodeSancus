package database

import (
	"fmt"
	"log"
	"os"
	"project/CodeSancus-Backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connects to the database
func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database!")
	} else {
		log.Println("Connected to database successfully")
	}

	db.AutoMigrate(&models.ActiveStorageAttachments{})
	db.AutoMigrate(&models.ActiveStorageBlobs{})
	db.AutoMigrate(&models.ActiveStorageVariantRecords{})
	db.AutoMigrate(&models.Datasets{})
	db.AutoMigrate(&models.DelayedJobs{})
	db.AutoMigrate(&models.Reports{})

	DB = db
}
