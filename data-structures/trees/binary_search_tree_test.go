package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *BinaryTree
		input    int
		expected *BinaryTree
	}{
		"inserts to empty tree": {
			setup:    NewBinaryTree,
			input:    10,
			expected: &BinaryTree{root: &node{value: 10, height: 1}},
		},
		"inserts smaller element": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)

				return tree
			},
			input:    1,
			expected: &BinaryTree{root: &node{value: 10, height: 3, left: &node{value: 5, height: 2, left: &node{value: 1, height: 1}}}},
		},
		"inserts larger element": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(15)

				return tree
			},
			input:    20,
			expected: &BinaryTree{root: &node{value: 10, height: 3, right: &node{value: 15, height: 2, right: &node{value: 20, height: 1}}}},
		},
		"insert duplicate element": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)

				return tree
			},
			input:    10,
			expected: &BinaryTree{root: &node{value: 10, height: 1}},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := testCase.setup()

			tree.Insert(testCase.input)

			assert.Equal(t, testCase.expected, tree)
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *BinaryTree
		input    int
		expected bool
	}{
		"searches for root": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    10,
			expected: true,
		},
		"searches for leaf": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    5,
			expected: true,
		},
		"searches for non-leaf element": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    15,
			expected: true,
		},
		"searches for element not in tree to the left": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    1,
			expected: false,
		},
		"searches for element not in tree to the right": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    20,
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := testCase.setup()

			assert.Equal(t, testCase.expected, tree.Search(testCase.input))
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *BinaryTree
		input    int
		expected *BinaryTree
	}{
		"deletes left leaf": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    5,
			expected: &BinaryTree{root: &node{value: 10, height: 2, right: &node{value: 15, height: 1}}},
		},
		"deletes right leaf": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)

				return tree
			},
			input:    15,
			expected: &BinaryTree{root: &node{value: 10, height: 2, left: &node{value: 5, height: 1}}},
		},
		"deletes non-leaf element": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    15,
			expected: &BinaryTree{root: &node{value: 10, height: 3, left: &node{value: 5, height: 1}, right: &node{value: 12, height: 2, right: &node{value: 17, height: 1}}}},
		},
		"deletes element with left child only": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)

				return tree
			},
			input:    15,
			expected: &BinaryTree{root: &node{value: 10, height: 2, left: &node{value: 5, height: 1}, right: &node{value: 12, height: 1}}},
		},
		"deletes element with right child only": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(17)

				return tree
			},
			input:    15,
			expected: &BinaryTree{root: &node{value: 10, height: 2, left: &node{value: 5, height: 1}, right: &node{value: 17, height: 1}}},
		},
		"deletes root": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(7)
				tree.Insert(9)
				tree.Insert(15)

				return tree
			},
			input:    10,
			expected: &BinaryTree{root: &node{value: 9, height: 3, left: &node{value: 5, height: 2, right: &node{value: 7, height: 1}}, right: &node{value: 15, height: 1}}},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := testCase.setup()

			tree.Delete(testCase.input)

			assert.Equal(t, testCase.expected, tree)
		})
	}
}

func TestTraverse(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *BinaryTree
		input    TraversalMethod
		expected []int
	}{
		"traverses in-order": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    InOrder,
			expected: []int{5, 10, 12, 15, 17},
		},
		"traverses pre-order": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    PreOrder,
			expected: []int{10, 5, 15, 12, 17},
		},
		"traverses post-order": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    PostOrder,
			expected: []int{5, 12, 17, 15, 10},
		},
		"traverses pre-order by default": {
			setup: func() *BinaryTree {
				tree := NewBinaryTree()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(17)

				return tree
			},
			input:    TraversalMethod(""),
			expected: []int{5, 10, 12, 15, 17},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := testCase.setup()

			values := tree.Traverse(testCase.input)

			assert.Equal(t, testCase.expected, values)
		})
	}
}
