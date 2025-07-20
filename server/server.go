package server

/* //go:generate go-bindata-assetfs -o web_static.go web/dist/... */

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/server/api"
)

func main() {
	server := gin.Default()
	// server.Use(cors.Default())

	fmt.Println("Starting server...")
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	fmt.Println("Setting up API routes...")

	api.SetupRoutes(server)
	fmt.Println("API routes set up successfully.")

	// Register catch-all route AFTER api routes
	/* server.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html") // or your static file handler
	}) */

	server.Static("/", "./web/dist")

	/* staticFS, err := fs.Sub(web.StaticFiles, "dist")
	if err != nil {
		log.Fatalf("failed to load embedded files: %v", err)
	}
	server.StaticFS("/view/", http.FS(staticFS)) */

	// Fallback to index.html for SPA routes
	/* server.NoRoute(func(c *gin.Context) {
		c.FileFromFS("./web/dist/index.html", http.FS(staticFS))
	}) */

	log.Println("🚀 Server listening at http://localhost:8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func Start() {
	main() // wraps gin setup
}
