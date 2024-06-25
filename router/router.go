package routes

import (
	"event_scheduler/handler"
	"event_scheduler/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.Login)
	r.GET("/events", middleware.IsLogin(), handler.FindEvent)
	return r
}
