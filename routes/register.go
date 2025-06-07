package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId) // Check if event exists
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered!", "event": event})
}

func unregisterFromEvent(context *gin.Context) {}
