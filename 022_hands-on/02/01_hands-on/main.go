package main

import (
	"io"
	"net"
)

func main() {

	port := ":8080"

	// The server should use net.Listen to listen on port 8080.
	// Remember to close the listener using defer.
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.
	// Listen and accept connections on port 8080 using net.Listen and net.Conn.
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// Write a response back on the connection.
		//Use io.WriteString to write the response: I see you connected.
		io.WriteString(conn, "I see you connected.")
		conn.Close()
	}

}
