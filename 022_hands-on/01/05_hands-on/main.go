package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(Slash))
	http.Handle("/dog/", http.HandlerFunc(Dog))
	http.Handle("/me/", http.HandlerFunc(Me))

	//ListenAndServe on port ":8080" using the default ServeMux
	http.ListenAndServe(":8080", nil)
}

func Slash(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("/Users/zmey56/Projects/golang-web-dev/022_hands-on/01/03_hands-on/templates/index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	//a template is parsed and served
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Index")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func Dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("/Users/zmey56/Projects/golang-web-dev/022_hands-on/01/03_hands-on/templates/dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	//a template is parsed and served
	err = tpl.ExecuteTemplate(res, "dog.gohtml", "Dog")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func Me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("/Users/zmey56/Projects/golang-web-dev/022_hands-on/01/03_hands-on/templates/me.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	//a template is parsed and served
	err = tpl.ExecuteTemplate(res, "me.gohtml", "Sasha")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
