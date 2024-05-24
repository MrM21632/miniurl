package main

import (
	"miniurl/miniurl"
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
	router.GET("/listall", func(c *gin.Context) {
		records, err := miniurl.GetAllRecords()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{"records": records})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Endpoint not found"})
	})

	router.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
