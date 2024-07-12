package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth2Callback this webpage is shown when user authorize himself for first time while creating an event
func Auth2Callback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You can exit this page"})
}
