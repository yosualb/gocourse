package main

import "errors"
import "fmt"

func RaiseError() error {
	// Error
	return errors.New("This is error")
}

func main() {
	err := RaiseError()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
