package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/thijsoostdam/kitecrewserver/model"
)

type GeoLocation struct {
	Latitude  float64
	Longitude float64
}

type RegistrationBody struct {
	Code     string      `json:"code"`
	AppID    string      `json:"app-id"`
	Location GeoLocation `json:"location"`
}
type InviteBody struct {
	Code     string      `json:"code"`
	AppID    string      `json:"app-id"`
	Location GeoLocation `json:"location"`
}

func UserIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(model.GetUsers()); err != nil {
		panic(err)
	}
}

func UserGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	log.Printf(
		"%s\t%s\t%s\t%s",
		"getting user:",
		username)
		
	user := model.GetUser(username)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	var regBody RegistrationBody
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &regBody); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//TODO: security, check app-id and stuffs.
	model.User.Register(username, regBody.Code)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Invite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var invBody InviteBody
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &invBody); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)

		}
	}

	invite := model.Invite{username, invBody.Code}
	data.StoreInvite(invite)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(invite); err != nil {
		panic(err)
	}
}
