package main

import (
	db "event_scheduler/database"
	"event_scheduler/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.Login)
	r.Run(":8000")
}
