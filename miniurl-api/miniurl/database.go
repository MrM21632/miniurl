package miniurl

import (
	"errors"
	"log"
	"miniurl/miniurl/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.URL{})
	Database = database
}

func WriteNewRecordToDatabase(long_url string, short_url string) (*models.URL, error) {
	new_record := models.URL{OriginalURL: long_url, ShortenedURL: short_url}
	result := Database.Create(&new_record)
	if result.RowsAffected < 1 {
		return nil, errors.New("failed to save new record to database")
	}
	if result.Error != nil {
		log.Println("Error occurred during write: " + result.Error.Error())
		return nil, result.Error
	}

	return &new_record, nil
}
