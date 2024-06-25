package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db := "root:root@/event_scheduler"
	var err error
	DB, err = sql.Open("mysql", db)
	if err != nil {
		log.Fatal("Failed to connect to db \n", err)
		return
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to Ping \n", err)
		return
	}
	fmt.Println("hello")
}
