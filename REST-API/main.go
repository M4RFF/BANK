package main

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // creates a server

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.POST("/events", postEvent)

	server.Run(":1313") // localhost 1313
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func postEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // kind of similar with fmt.Scan(&...)

	if err != nil {
		// respond to client that smth went wrong (400).
		// gin.H() creates a map that display a message
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse  request data"})
		// if it went wrong we need to return
		return
	}

	event.ID = 1
	event.UserID = 1
	// send a respond if everything works
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
