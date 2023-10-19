package main

import "net/http"

//ListenAndServe on port ":8080" using the default ServeMux.
//
//Use HandleFunc to add the following routes to the default ServeMux:
//
//"/"
//"/dog/"
//"/me/
//

//
//Have the "/me/" route print out your name.

func main() {
	//Use HandleFunc to add the following routes to the default ServeMux
	http.HandleFunc("/", Slash)
	http.HandleFunc("/dog/", Dog)
	http.HandleFunc("/me/", Me)

	//ListenAndServe on port ":8080" using the default ServeMux
	http.ListenAndServe(":8080", nil)
}

// Add a func for each of the routes.
// Slash function for handlefunc to handle the "/" route
func Slash(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Slash"))
}

// Dog function for handlefunc to handle the "/dog/" route
func Dog(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Dog"))
}

// Me function for handlefunc to handle the "/me/" route
// Have the "/me/" route print out your name.
func Me(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Sasha"))
}
