package string

import (
	"bytes"
)

// Unexported function
func addNamePrefix(s string) string {
	var buff bytes.Buffer

	buff.WriteString(s)
	buff.WriteString(Name)

	return buff.String()
}
