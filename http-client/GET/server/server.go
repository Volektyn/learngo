package main

import (
	"log"
	"net/http"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// msg := "{\"message\": \"hello world\"}"
	msg := "{\"names\": [\"Electric\", \"Boogaloo\", \"Electric\", \"Boogaloo\", \"Boogaloo\", \"Electric\", \"Boogaloo\", \"Boogaloo\"]}"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
