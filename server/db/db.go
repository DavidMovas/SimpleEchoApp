package db

import (
	"echoapp/entities"
	"log"

	"go.etcd.io/bbolt"
)

type DB struct {
	*bbolt.DB
}

func NewDB() (*DB, error) {
	db, err := bbolt.Open("database.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var DB = &DB{db}

	err = DB.initDB()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return DB, nil
}

func (db *DB) initDB() error {
	_, err := db.createNewBucket("users")

	var users []entities.User

	users = append(users, *entities.NewUser("David", 21, "david417@gmail.com"))
	users = append(users, *entities.NewUser("Dima", 20, "dima2003@gmail.com"))
	users = append(users, *entities.NewUser("Nika", 25, "nK6Bh@example.com"))
	users = append(users, *entities.NewUser("Jane", 30, "jane@ex.com"))
	users = append(users, *entities.NewUser("Jim", 32, "jim@ex.com"))
	users = append(users, *entities.NewUser("Dom", 28, "dom522@ex.com"))
	users = append(users, *entities.NewUser("Max", 35, "msx@ex.com"))

	for _, u := range users {
		_ = db.AddUser(&u)
	}

	return err
}

func (db *DB) createNewBucket(bucketName string) (bucket *bbolt.Bucket, err error) {

	err = db.Update(func(tx *bbolt.Tx) error {
		bucket, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	if err != nil {
		return nil, err
	}

	return bucket, err
}
