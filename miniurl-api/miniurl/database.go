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

func CheckForLongUrlInDatabase(long_url string) (bool, error) {
	found := models.URL{}
	result := Database.Where("original_url = ?", long_url).First(&found)
	if result.Error != nil {
		log.Println("Error occurred during query: " + result.Error.Error())
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func CheckForShortUrlInDatabase(short_url string) (bool, error) {
	found := models.URL{}
	result := Database.Where("shortened_url = ?", short_url).First(&found)
	if result.Error != nil {
		log.Println("Error occurred during query: " + result.Error.Error())
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func GetRecord(shortened_url string) (*models.URL, error) {
	var record models.URL
	result := Database.Where("shortened_url = ?", shortened_url).First(&record)
	if result.Error != nil {
		log.Println("Error occurred during query: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found in database")
	}

	return &record, nil
}

func GetAllRecords() ([]models.URL, error) {
	var records []models.URL
	result := Database.Find(&records)
	if result.Error != nil {
		log.Println("Error occurred during read: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no records found")
	}

	return records, nil
}
