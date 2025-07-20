package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ndkariyasena/gobolt-view/server/db"
)

func SetupRoutes(server *gin.Engine) {
	router := server.Group("/api")

	router.GET("/connection/is-active", db.IsDatabaseConnected)
	router.POST("/config/db", db.LoadDatabaseHandler)
	router.GET("/config/db", db.GetCurrentDatabaseHandler)
	router.GET("/buckets", db.ListBucketsHandler)
	router.GET("/bucket/:name/keys", db.ListKeysHandler)
	router.GET("/bucket/:name/key-value", db.ListKeyValuesHandler)
	router.GET("/bucket/:name/key/:key", db.GetValueHandler)
}
