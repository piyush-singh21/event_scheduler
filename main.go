package main

import (
	db "event_scheduler/database"
	routes "event_scheduler/router"

	_ "event_scheduler/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Documenting API (Event Scheduling)
// @Version 1.0
// @Description Event Schedular
// @contact.name Piyush Singh
// @contact.url piyush.singh@stl.tech
// @contact.email singh.piyush00034@gmail.com
// @Host localhost:8000
func main() {
	db.InitDB()
	r := routes.SetupRoute()
	// ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
