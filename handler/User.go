package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser registers a new user to the database
// @Summary Register a new user
// @Description Register a new user and add the details in the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body User true "User registration request"
// @Success 200 {object} User
// @Failure 400 {object} User
// @Failure 500 {object} User
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.ValidateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
