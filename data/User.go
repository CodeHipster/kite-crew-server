package data

//TODO: make this an interface so we can use something else at some other time.

import (
	"fmt"

	"github.com/thijsoostdam/kitecrewserver/model"
)

var currentUserId int

var users model.Users

type Users []User

// Give us some seed data
func init() {
	StoreUser(model.User{Username: "Kees"})
	StoreUser(model.User{Username: "Gertrude"})
}

func GetUsers() []User {
	return users
}

func FindUser(username string) *model.User {
	for _, u := range users {
		if u.Username == username {
			return &u
		}
	}
	// return nil if not found
	return nil
}

func StoreUser(u model.User) error {
	if FindUser(u.Username) != nil {
		return fmt.Errorf("User with Username: %s already exists", u.Username)
	}
	currentUserId += 1
	users = append(users, u)
	return nil
}

func DestroyUser(username string) error {
	for i, u := range users {
		if u.Username == username {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with username %s to delete", username)
}
