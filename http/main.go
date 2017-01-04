package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World!"))
}

func main() {
	// HTTP server
	go func() {
		http.HandleFunc("/hello", home)
		// HTTP server listens to selected port
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}()

	// HTTP client
	go func() {
		// HTTP client sends request to server
		res, err := http.Get("http://127.0.0.1:8080/hello")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		// Read response
		bs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(string(bs))
	}()

	var input string
	fmt.Scanln(&input)
}
