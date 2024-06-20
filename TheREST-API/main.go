package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// configers an http server
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, DELETE, PUT, PATCH

	// start it and starts listening for incomming requests
	server.Run(":1331") // loacalhost: 1331
}

func getEvents(context *gin.Context) {
	// send back a respond by JSON function

	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) // stores tthe number 200
}
