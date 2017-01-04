package string

import "testing"

// Test case struct
type TestCase struct {
	input    interface{}
	expected interface{}
}

// Test cases
var addPrefixTestCases = []TestCase{
	{"Yosua", "Mr. Yosua"},
	{"John", "Mr. John"},
	{"", "Mr. "},
	{nil, "Mr. "},
}

// Testing function
func TestAddPrefix(t *testing.T) {
	for _, tc := range addPrefixTestCases {
		if res := AddPrefix(tc.input.(string)); res != tc.expected.(string) {
			t.Error("Expected:", tc.expected.(string), " Got:", res)
		}
	}
}
