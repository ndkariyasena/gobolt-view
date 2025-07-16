package db

import (
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func InitDB(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, nil)
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
