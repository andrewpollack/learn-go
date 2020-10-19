package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("%v\n", conn)
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Printf("ASDF\n")
		}
		go handleConnection(conn)
	}
}
