package strings

// Reverse uses two indices to swap characters outside-in from the start and end of the string.
// Though twice as fast as the slow solution, it is technically still O(n).
// Time: O(n)
// Space: O(n)
func Reverse(s string) string {
	n := len(s)
	chars := []rune(s)
	reversed := make([]rune, n)

	start := 0
	end := n - 1
	for start < n/2 {
		reversed[start], reversed[end] = chars[end], chars[start]
		start++
		end--
	}

	// if there's an odd number of chars, the middle char is stays the same
	if n%2 != 0 {
		reversed[n/2] = chars[n/2]
	}

	return string(reversed)
}

// SlowReverse iterates over the whole string.
// Time: O(n)
// Space: O(n)
func SlowReverse(s string) string {
	reversed := make([]rune, len(s))

	n := len(s) - 1

	for i, char := range s {
		reversed[n-i] = char
	}

	return string(reversed)
}
