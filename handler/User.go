package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Register user by taking his email and password
// @Schemes
// @Description register user
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.User true "Register User"
// @Success 200 {User created successfully}
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
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
