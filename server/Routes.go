package server

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"UserIndex", "GET", "/users", UserIndex},
	Route{"GetUser", "GET", "/users/{username}", UserGet},
	Route{"Register", "POST", "/users/{username}/register", Register},
	Route{"Invite", "POST", "/users/{username}/invite", Invite},
}
