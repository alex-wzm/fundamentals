package strings

import (
	"unicode"
)

// isPalindrome checks if a string is a palindrome
// O(n) time complexity
// O(n) space complexity (can be improved to O(1) at the cost of readability)
func isPalindrome(s string) bool {
	chars := []rune{}

	// filter spaces from string
	// and convert all characters to lowercase
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			chars = append(chars, unicode.ToLower(char))
		}
	}

	i, j := 0, len(chars)-1

	for i < j {
		if chars[i] != chars[j] {
			return false
		}

		i++
		j--
	}

	return true
}
