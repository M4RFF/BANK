package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // creates a server

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.Run(":1313") // localhost 1313
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
