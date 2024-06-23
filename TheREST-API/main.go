package main

import (
	"net/http"
	"theRestApi/db"
	"theRestApi/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	// configers an http server
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, DELETE, PUT, PATCH

	server.POST("/events", postEvent) // the same path /events

	// start it and starts listening for incomming requests
	server.Run(":1331") // loacalhost: 1331
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// send back a respond by JSON function
	context.JSON(http.StatusOK, events) // stores the number 200
}

func postEvent(context *gin.Context) { // we use again (context *gin.Context)
	// we need to get some data of incoming requests
	var event models.Event
	// if some filds are missing it's not a problem, because line bellow
	// just set a missing filed as zero
	err := context.ShouldBindJSON(&event) // a little bit works like a Scan()

	if err != nil {
		// it send back a respond that tells a user that smth went wrong
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
