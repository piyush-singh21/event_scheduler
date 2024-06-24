package routes

import (
	"event_scheduler/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.Login)
	return r
}
