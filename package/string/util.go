package string

import (
	"bytes"
)

// Unexported function
func addNamePrefix(s string) string {
	var buff bytes.Buffer

	if _, err := buff.WriteString(s); err != nil {
		return buff.String()
	}
	if _, err := buff.WriteString(Name); err != nil {
		return buff.String()
	}

	return buff.String()
}
