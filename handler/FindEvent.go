package handler

import (
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindEvent(c *gin.Context) {
	res, err := service.GetAllEvent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
