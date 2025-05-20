package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torcuata22/rest_api/db"
	"github.com/torcuata22/rest_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
