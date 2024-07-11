package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth2Callback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You can exit this page"})
}
