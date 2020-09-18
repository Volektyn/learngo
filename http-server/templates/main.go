package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type User struct {
	Username string
}

type WelcomeTmpl struct {
	tmpl *template.Template
}

func (tmpl WelcomeTmpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	un := r.URL.Query()
	user := User{}
	username, ok := un["username"]
	if ok {
		user.Username = strings.Join(username, ",")
	} else {
		user.Username = "visitor"
	}
	tmpl.tmpl.Execute(w, user)
}

func NewWelcome(tmplPath string) (*WelcomeTmpl, error) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}
	return &WelcomeTmpl{tmpl}, nil
}

func main() {
	welcome, err := NewWelcome("index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
