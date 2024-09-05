package db

import (
	"encoding/json"
	"errors"

	. "echoapp/entities"
	"go.etcd.io/bbolt"
)

func (db *DB) AddUser(user *User) (err error) {
	err = db.Update(func(tx *bbolt.Tx) error {

		b := tx.Bucket([]byte("users"))

		var userByte []byte

		userByte, err = json.Marshal(user)
		err = b.Put([]byte(user.Email), userByte)

		return err
	})

	return err
}

func (db *DB) UpdateUser(user *User) (err error) {
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		isUser := (b.Get([]byte(user.Email))) != nil

		if !isUser {
			return errors.New("no user found")
		}

		err = b.Delete([]byte(user.Email))

		var userByte []byte

		userByte, err = json.Marshal(user)
		err = b.Put([]byte(user.Email), userByte)

		return err
	})

	return err
}

func (db *DB) GetUsers() (users []User, err error) {
	users = make([]User, 0)

	err = db.View(func(tx *bbolt.Tx) error {

		b := tx.Bucket([]byte("users"))

		err = b.ForEach(func(k, v []byte) error {
			var user User
			err = json.Unmarshal(v, &user)

			users = append(users, user)

			return err
		})

		return err
	})

	return users, err
}

func (db *DB) GetUser(email string) (user *User, err error) {
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		userByte := b.Get([]byte(email))

		if userByte == nil {
			return errors.New("no user found")
		}

		err = json.Unmarshal(userByte, &user)

		return err
	})

	return user, err
}

func (db *DB) DeleteUser(email string) (err error) {
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		err = b.Delete([]byte(email))

		return err
	})

	return err
}
