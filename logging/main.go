package main

import (
	"errors"
	"log"
	"os"
)

func RaiseError() error {
	return errors.New("This is error.")
}

func main() {
	err := RaiseError()
	if err != nil {
		// Standard output logging
		log.Println(err)

		// File logging
		f, _ := os.Create("logging/log")
		log.SetOutput(f)
		log.Println(err)
	}

}
