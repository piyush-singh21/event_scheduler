package service

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	auth "event_scheduler/utils"

	"github.com/gin-gonic/gin"
)

func ValidateRegistrationId(id int, c *gin.Context) (model.EventResp, error) {
	var eventId int
	var email string
	var registerEmail string
	database.DB.QueryRow("SELECT id from events WHERE id=?", id).Scan(&eventId)
	if id != eventId {
		return model.EventResp{}, errors.New("event does not exist")
	}
	keyInterface, _ := c.Get("id")
	key := auth.Convert(keyInterface)
	eventResp := database.FindEvent(id)
	database.DB.QueryRow("SELECT email from users WHERE id=?", key).Scan(&email)
	database.DB.QueryRow("SELECT email from register WHERE email=? AND eventId=?", email, id).Scan(&registerEmail)
	if registerEmail == email {
		return model.EventResp{}, errors.New("already registered to event")
	}

	database.RegisterForEvent(eventId, key)
	return eventResp, nil

}
