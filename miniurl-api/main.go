package main

import (
	"miniurl/miniurl"
	"miniurl/miniurl/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	miniurl.ConnectToDatabase()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/shorten", miniurl.CreateUrlRecord)
	router.GET("/navigate", func(c *gin.Context) {
		var input models.NavigateURLInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		record, err := miniurl.GetRecord(input.ShortenedURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
			return
		}
		c.Redirect(http.StatusFound, record.OriginalURL)
	})
	router.GET("/listall", func(c *gin.Context) {
		records, err := miniurl.GetAllRecords()
		if err != nil || len(records) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"records": records})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Endpoint not found"})
	})

	router.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
