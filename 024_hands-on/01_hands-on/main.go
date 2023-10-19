package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

// For the default route "/"  Have a func called "foo"  which writes to the response "foo ran"
func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

// For the route "/dog/"  Have a func called "dog"  which parses a template called "dog.gohtml" and writes to
// the response "<h1>This is from dog</h1>"  and also shows a picture of a dog when the template is executed.
func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("/Users/zmey56/Projects/golang-web-dev/024_hands-on/01_hands-on/dog.gohtml")
	if err != nil {
		log.Fatal("Problem with ParseFiles")
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatal("Problem with ExecuteTemplate")
	}
}

// Use "http.ServeFile"  to serve the file "dog.jpeg"
func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/zmey56/Projects/golang-web-dev/024_hands-on/01_hands-on/toby.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dogpic", dogPic)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
