package main

import (
	"miniurl/miniurl"
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
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Endpoint not found"})
	})

	router.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
