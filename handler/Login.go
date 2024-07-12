package handler

import (
	"event_scheduler/model"
	auth "event_scheduler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login user by taking his email and password
// @Schemes
// @Description register user
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.LoginRequest true "Login User"
// @Success 200 {string} token
// @Router /login [post]

// Login Registered users can login, once you login you will be provided with a token
func Login(c *gin.Context) {
	var logReq model.LoginRequest
	if err := c.ShouldBindJSON(&logReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.AuthenticateUser(logReq.Email, logReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
