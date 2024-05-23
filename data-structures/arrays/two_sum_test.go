package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		input    []int
		target   int
		expected []int
	}{
		"Two numbers sum to target": {
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 0},
		},
		"Three numbers sum to target": {
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{2, 1},
		},
		"Two identical numbers sum to target": {
			input:    []int{3, 3},
			target:   6,
			expected: []int{1, 0},
		},
		"Target not found": {
			input:    []int{1, 2, 3, 4, 5},
			target:   10,
			expected: []int{},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			if got := twoSum(testCase.input, testCase.target); !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("twoSum() = %v, want %v", got, testCase.expected)
			}
		})
	}
}
