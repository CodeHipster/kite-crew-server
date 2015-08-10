package model

import (
	"errors"
	"fmt"
	"github.com/thijsoostdam/kitecrewserver/data"
)

//TODO: make modeltypes only visible to model package.

type User struct {
	Username string `json:"username"`
}


func GetUser(username string) (User, error) {
	user := data.FindUser(username)
	if user == nil {
		return User{}, errors.New(fmt.Errorf("could not find user: %s", username))
	}
	return user, nil
}

func GetUsers() []User{
	return data.GetUsers()
}

func (user User) Register(username string, code string) User{

//TODO: check if code is in repo.
	user := User{username}
	data.StoreUser(user)
	return User
}