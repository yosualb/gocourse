package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Server struct
type Server struct{}

// Server method
func (s *Server) AddPrefix(str string, res *string) error {
	*res = "Mr. " + str
	return nil
}

// RPC server
func server() {
	// Register RPC server
	rpc.Register(new(Server))
	// RPC server listens to selected protocol and port
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for {
		// RPC server accepts incoming connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		// RPC server handles connection concurrently
		go rpc.ServeConn(c)
	}
}

// RPC client
func client() {
	// RPC client sends request to server
	c, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var res string
	// RPC client calls server method
	err = c.Call("Server.AddPrefix", "Yosua", &res)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Server.AddPrefix(\"Yosua\"):", res)
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
