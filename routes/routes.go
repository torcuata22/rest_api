package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")          //creates a new group
	authenticated.Use(middlewares.Authenticate) //always runs middleware before handlers
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}

//Individual implementation of middleware and auth:
//	server.POST("/events", middlewares.Authenticate, createEvent) //protected and runs Authenticate() midleware BEFORE the handler
