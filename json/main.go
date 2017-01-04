package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// JSON struct
type Person struct {
	Name string
	// JSON alias
	Sex string `json:"gender"`
}

func main() {
	// JSON Marshal
	p := &Person{"Yosua", "M"}
	b, _ := json.Marshal(p)
	fmt.Println("string(b):", string(b))

	// JSON Encoder
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(&p)

	// JSON Unmarshal
	data := []byte(`{"Name": "Yosua", "Gender": "M"}`)
	var p1 *Person
	json.Unmarshal(data, &p1)
	fmt.Println("p1:", p1)

	// JSON Decoder
	dec := json.NewDecoder(strings.NewReader(string(data)))
	dec.Decode(&p1)
	fmt.Println("p1:", p1)
}
