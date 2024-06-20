package service

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	"regexp"
	"strings"
)

func ValidateUser(user model.User) error {
	if len(user.Name) <= 2 {
		return errors.New("length of string should be greater than 2")

	}
	if match, _ := regexp.MatchString("^[a-zA-Z]+$", user.Name); !match {
		return errors.New("not valid name")

	}
	if !strings.ContainsAny(user.Email, "@") || len(user.Email) < 5 {
		return errors.New("incorrect email")
	}
	err := database.AddUser(user)
	if err != nil {
		return err
	}

	return nil

}
