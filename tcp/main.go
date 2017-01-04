package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// TCP server
func server() {
	// TCP server listens to selected protocol and port
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for {
		// TCP server accepts incoming connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		// TCP server handles request concurrently
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	var msg string
	// Decode message from request
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Received", msg)

	// Close connection
	err = c.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

// TCP client
func client() {
	// TCP client sends request to server
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	msg := "Hello World!"
	fmt.Println("Sending", msg)
	// Encode message
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("err:", err)
	}

	// Close connection
	err = c.Close()
	if err != nil {
		fmt.Println("err:", err)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
