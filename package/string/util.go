package string

import (
	"bytes"
)

// Unexported function
func addPrefix(s string) string {
	var buff bytes.Buffer

	buff.WriteString("Prefix + ")
	buff.WriteString(s)

	return buff.String()
}
