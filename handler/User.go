package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegResp struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
