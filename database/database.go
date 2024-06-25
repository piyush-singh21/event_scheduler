package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
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
		createTable()
	})
}
func createTable() {
	userTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(100) UNIQUE NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL
		);
	`
	eventTable := `
		CREATE TABLE IF NOT EXISTS events(
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			description TEXT,
			date DATETIME NOT NULL
		);
	`
	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal("Failed to create user Table", err)
		return
	}
	_, err = DB.Exec(eventTable)
	if err != nil {
		log.Fatal("Failed to create user Table", err)
		return
	}

}
