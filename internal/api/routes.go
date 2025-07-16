package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/internal/db"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/config/db", db.LoadDatabaseHandler)
	r.GET("/config/db", db.GetCurrentDatabaseHandler)
	r.GET("/buckets", db.ListBucketsHandler)
	r.GET("/bucket/:name/keys", db.ListKeysHandler)
	r.GET("/bucket/:name/key/:key", db.GetValueHandler)
}
