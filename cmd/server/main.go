package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/internal/api"
)

func main() {
	r := gin.Default()

	api.SetupRoutes(r)

	log.Println("🚀 Server listening at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
