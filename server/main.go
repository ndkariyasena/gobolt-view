package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/server/api"
)

var portNumber string

func main() {
	flag.StringVar(&portNumber, "port", "8080", "Port number for the server to listen on")
	flag.Parse()

	server := gin.Default()
	server.Use(cors.Default())

	fmt.Println(portNumber)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api.SetupRoutes(server)

	log.Printf("🚀 Server listening at http://localhost:%v", portNumber)

	if err := server.Run(":" + portNumber); err != nil {
		log.Fatal(err)
	}
}
