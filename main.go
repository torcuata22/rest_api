package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/db"
	"github.com/torcuata22/rest_api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get events"})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get event"})
	}
	context.JSON(http.StatusOK, event)
}
