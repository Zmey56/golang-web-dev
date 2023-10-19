package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	port := ":8080"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)

	}

}

// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
// Pass the connection of type net.Conn as an argument into this function.
func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("This is the end of the http request headers.")
			break
		}
	}
}
