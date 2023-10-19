package main

import (
	"log"
	"net/http"
)

//# Serve the files in the "starting-files" folder

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
