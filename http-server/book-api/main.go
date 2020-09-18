package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Counter int    `json:"views"`
}

func (book *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}
	book.Counter++
	data, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func main() {
	mainPage := PageWithCounter{Content: "This is main page", Title: "Hello World!"}
	cha1 := PageWithCounter{Content: "This is chapter 1", Title: "Chapter 1"}
	cha2 := PageWithCounter{Content: "This is chapter 2", Title: "Chapter 2"}

	http.Handle("/", &mainPage)
	http.Handle("/cha1", &cha1)
	http.Handle("/cha2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
