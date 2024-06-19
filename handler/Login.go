package handler

import (
	"event_scheduler/auth"
	"event_scheduler/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var logReq model.LoginRequest
	if err := c.ShouldBindJSON(&logReq); err != nil {
		fmt.Println(15)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.AuthenticateUser(logReq.Email, logReq.Password)
	if err != nil {
		fmt.Println(20)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
