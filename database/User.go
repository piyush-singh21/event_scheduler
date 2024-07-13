package database

import (
	"event_scheduler/model"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(user model.User) error {
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err := DB.Exec("INSERT INTO users(username,email,password) VALUES(?,?,?)", user.Name, user.Email, string(pass))
	if err != nil {
		return err
	}
	return nil
}

func GetUserMail(id int) string {
	var mail string
	DB.QueryRow("SELECT email FROM users WHERE id=?", id).Scan(&mail)
	return mail
}
