package main

import (
	db "event_scheduler/database"
	routes "event_scheduler/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.InitDB()
	r := routes.SetupRoute()
	// ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
