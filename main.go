package main

import (
	"log"
	"net/http"
	"github.com/leitu/roots/config"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It Works!"))
}

func main() {
	http.Handle("/", greeting)
	log.Println("Starting IPXEServer")
	log.Fatal(http.ListenAndServe(config.ListernAddr, nil))
}