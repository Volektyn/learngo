// 4. Define a struct called PageWithCounter with a counter as an integer attribute, a
// content, and a heading as a text attribute.
// 5. Add a ServeHTTP method to the struct, capable of displaying the content, the
// heading, and a message with the total number of views.
// 6. Create your main function and, inside, implement the following:
// 7. Instantiate three handlers of the PageWithCounter type, with Hello World, Chapter 1,
// and Chapter 2 headings and some content.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (pwc *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}
	pwc.counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<br/><div><p>%s</p>\n<p>Page counter: %d</p></div>", pwc.heading, pwc.content, pwc.counter)
	w.Write([]byte(msg))
}
func main() {
	mainPage := PageWithCounter{content: "This is main page", heading: "Hello World!"}
	cha1 := PageWithCounter{content: "This is chapter 1", heading: "Chapter 1"}
	cha2 := PageWithCounter{content: "This is chapter 2", heading: "Chapter 2"}

	http.Handle("/", &mainPage)
	http.Handle("/cha1", &cha1)
	http.Handle("/cha2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
