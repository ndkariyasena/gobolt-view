package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

func LoadDatabaseHandler(c *gin.Context) {
	var req struct {
		Path string `json:"dbFilePath"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := LoadDatabase(req.Path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Database loaded", "path": req.Path})
}

func GetCurrentDatabaseHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"path": GetCurrentPath()})
}

func ListBucketsHandler(c *gin.Context) {
	var buckets []string

	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buckets": buckets})
}

func ListKeysHandler(c *gin.Context) {
	bucket := c.Param("name")
	var keys []string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, _ []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"keys": keys})
}

func GetValueHandler(c *gin.Context) {
	bucket := c.Param("name")
	key := c.Param("key")

	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		value = b.Get([]byte(key))
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if value == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": string(value)})
}
