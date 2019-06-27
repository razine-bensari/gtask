package db

import (
	"time"

	"github.com/boltdb/bolt"
)

//A bucket is a key value pair
var gTaskBucket = []byte("gtasks")
var db *bolt.DB

//Initialize the db by opening it.
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(gTaskBucket)
		return err
	})
}
