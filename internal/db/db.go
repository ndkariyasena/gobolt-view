package db

import (
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func InitDB(path string) error {
	var err error
	log.Println("Starting to make the database connection.")
	db, err = bolt.Open(path, 0666, &bolt.Options{Timeout: 5 * time.Second})
	log.Println("Database connection successfully established!")
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func GetDB() *bolt.DB {
	return db
}
