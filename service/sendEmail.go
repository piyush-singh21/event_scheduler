package service

import (
	"errors"
	"event_scheduler/database"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// Send an automated mail to the user who created the event
func GetDataToSendMail(userId int, body, signal string) error {
	email := database.GetEmail(userId)
	subject := ""
	if signal == "register" {
		subject = "Registerd to event"

	} else {
		subject = "Event Created"
	}
	err := sendMail(email, subject, body)
	if err != nil {
		return fmt.Errorf("failed to send mail %v", err)
	}
	return nil
}
func sendMail(to string, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("unable to get .env file")
	}
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	if from == "" || password == "" {
		return fmt.Errorf("email address or password variable not set")
	}
	host := "smtp.gmail.com"
	port := "587"
	hostaddress := host + ":" + port
	auth := smtp.PlainAuth("", from, password, host)
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n\n" + body
	err = smtp.SendMail(hostaddress, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("unable to send mail %s", err)
	}
	return nil

}
