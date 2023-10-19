package main

import (
	"bufio"
	"fmt"
	"io"
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
			panic(err)
		}
		defer conn.Close()

		//You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.
		//Use bufio.NewScanner() to read from the connection.
		scan := bufio.NewScanner(conn)
		for scan.Scan() {
			ln := scan.Text()
			println(ln)
		}
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
	}

}
