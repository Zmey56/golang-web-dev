// Using cookies, track how many times a user has been to your website domain.
// Print their number of visits.

package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", tmp)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func tmp(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "0",
		Path:  "/",
	})

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	//convert string to int
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	count++
	//convert int back to string
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)

	io.WriteString(w, c.Value)
}
