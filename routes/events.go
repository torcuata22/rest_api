package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/models"
)

func getEvents(context *gin.Context) {
	//check token is present
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get events"})
	}
	context.JSON(http.StatusOK, events)
}

// needs to be authenticated (check for Token), auth is now done in the routes (through a group)
func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event"})
		return
	}

	//retrieve user id from context
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	//check event before updating, compare userID to Token's userID
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find event id"})
		return
	}

	//event will only be updated if the user is the one who created it
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event by id"})
		return
	}

	//only event creator can delete it:
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
