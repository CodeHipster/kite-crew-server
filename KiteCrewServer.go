package main

import (
	"log"
	"net/http"
	"github.com/thijsoostdam/kitecrewserver/server"
)

func main() {
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
