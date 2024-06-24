package main

import (
	db "event_scheduler/database"
	routes "event_scheduler/router"
)

func main() {
	db.InitDB()
	r := routes.SetupRoute()
	r.Run(":8000")
}
