package service

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"

	"github.com/gin-gonic/gin"
)

func ValidateRegistrationId(key, id int, c *gin.Context) (model.EventResp, error) {
	eventId := database.GetEventId(id)
	if id != eventId {
		return model.EventResp{}, errors.New("event does not exist")
	}
	eventResp := database.FindEvent(id)
	email := database.GetEmailFromUsers(key)
	registerEmail := database.GetEmailFromRegister(email, id)
	userId := database.GetUidFromEvents(eventId)
	if registerEmail == email || userId == key {
		return model.EventResp{}, errors.New("already registered to event")
	}

	database.RegisterForEvent(eventId, key)
	return eventResp, nil

}
