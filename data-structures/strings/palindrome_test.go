package strings

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected bool
	}{
		"empty string": {
			input:    "",
			expected: true,
		},
		"single character": {
			input:    "a",
			expected: true,
		},
		"two different characters": {
			input:    "ab",
			expected: false,
		},
		"two same characters": {
			input:    "aa",
			expected: true,
		},
		"palindrome with odd length": {
			input:    "aba",
			expected: true,
		},
		"palindrome with even length": {
			input:    "abba",
			expected: true,
		},
		"non-palindrome with odd length": {
			input:    "abc",
			expected: false,
		},
		"non-palindrome with even length": {
			input:    "abcd",
			expected: false,
		},
		"palindrome with mixed case": {
			input:    "Aba",
			expected: true,
		},
		"non-palindrome with spaces": {
			input:    "race care",
			expected: false,
		},
		"non-palindrome with numbers": {
			input:    "12312",
			expected: false,
		},
		"palindrome with spaces and mixed case": {
			input:    "A man a plan a canal Panama",
			expected: true,
		},
		"palindrome with numbers and mixed case": {
			input:    "12321",
			expected: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := isPalindrome(testCase.input)
			if got != testCase.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", testCase.input, got, testCase.expected)
			}
		})
	}
}
