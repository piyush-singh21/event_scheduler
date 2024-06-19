package handler

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user.Name) <= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Length of name should be greater than 2"})
		return
	}
	if match, _ := regexp.MatchString("^[a-zA-Z]+$", user.Name); !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot have special character or number"})
		return

	}
	if !strings.ContainsAny(user.Email, "@") || len(user.Email) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email"})
		return
	}
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err := database.DB.Exec("INSERT INTO users(username,email,password) VALUES(?,?,?)", user.Name, user.Email, string(pass))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	responseUser := model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": responseUser})
}
