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

		scan := bufio.NewScanner(conn)
		for scan.Scan() {
			ln := scan.Text()
			println(ln)
			if ln == "" {
				fmt.Println("This is the end of the http request headers.")
				break
			}
		}
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}

}
