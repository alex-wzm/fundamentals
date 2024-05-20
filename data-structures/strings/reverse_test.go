package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = map[string]struct {
	input    string
	expected string
}{
	"empty": {
		input:    "",
		expected: "",
	},
	"basic": {
		input:    "hello",
		expected: "olleh",
	},
	"spaced": {
		input:    "hello world",
		expected: "dlrow olleh",
	},
	"punctuation": {
		input:    "Hello, World!",
		expected: "!dlroW ,olleH",
	},
}

func TestReverse(t *testing.T) {
	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			reversed := Reverse(c.input)
			assert.Equal(t, c.expected, reversed)
		})
	}
}

func TestSlowReverse(t *testing.T) {
	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			reversed := SlowReverse(c.input)
			assert.Equal(t, c.expected, reversed)
		})
	}
}
