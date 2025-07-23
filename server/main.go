package main

import (
	"flag"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/server/api"
)

var portNumber string
var ginMode string

func main() {
	flag.StringVar(&portNumber, "port", "8000", "Port number for the server to listen on")
	flag.StringVar(&ginMode, "mode", "release", "Gin mode: debug, release, or test")
	flag.Parse()

	gin.SetMode(ginMode)
	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api.SetupRoutes(server)

	log.Printf("🚀 Server listening at http://localhost:%v", portNumber)

	if err := server.Run(":" + portNumber); err != nil {
		log.Fatal(err)
	}
}
