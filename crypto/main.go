package main

import (
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// SHA1
	s := sha1.New()
	s.Write([]byte("Hello World!"))
	bs := s.Sum([]byte{})
	fmt.Println("string(bs):", string(bs))

	// BCrypt
	// Generate BCrypt hash with selected cost
	bs, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Println(string(bs))

	// Compare BCrypt hash
	err = bcrypt.CompareHashAndPassword(bs, []byte("password"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// Check BCrypt hash cost
	cost, err := bcrypt.Cost(bs)
	if err != nil {
		return
	}
	fmt.Println("cost:", cost)
}
