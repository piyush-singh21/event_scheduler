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
