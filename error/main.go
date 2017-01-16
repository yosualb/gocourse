package main

import "errors"
import "fmt"

type CustomError struct {
	code int
}

func RaiseError() error {
	// Error
	return errors.New("This is error")
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Custom error: %v", e.code)
}

func NewCustomError(code int) error {
	return &CustomError{code}
}

func main() {
	err := RaiseError()
	if err != nil {
		fmt.Println("err:", err)
	}

	// Custom Error
	err2 := NewCustomError(404)
	if err2 != nil {
		fmt.Println("err2:", err2)
	}
}
