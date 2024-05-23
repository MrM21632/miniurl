package miniurl

import (
	"miniurl/miniurl/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUrlRecord(c *gin.Context) {
	var input models.CreateURLInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortened_url": ""})
}
