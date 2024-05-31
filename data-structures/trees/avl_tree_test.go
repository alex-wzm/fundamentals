package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalanceAVL(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *AVLTree
		expected *node
	}{
		"already balanced tree": {
			setup: func() *AVLTree {
				root := &node{value: 10, height: 2}
				root.left = &node{value: 5, height: 1}
				root.right = &node{value: 15, height: 1}

				return &AVLTree{
					root: root,
				}
			},
			expected: &node{
				value:  10,
				height: 2,
				left:   &node{value: 5, height: 1},
				right:  &node{value: 15, height: 1},
			},
		},
		"balances right-heavy tree with RR rotation": {
			setup: func() *AVLTree {
				root := &node{value: 10, height: 3}
				root.right = &node{value: 20, height: 2}
				root.right.right = &node{value: 30, height: 1}

				return &AVLTree{
					root: root,
				}
			},
			expected: &node{
				value:  20,
				height: 2,
				left:   &node{value: 10, height: 1},
				right:  &node{value: 30, height: 1},
			},
		},
		"balances right-heavy tree with RL rotation": {
			setup: func() *AVLTree {
				root := &node{value: 10, height: 3}
				root.right = &node{value: 30, height: 2}
				root.right.left = &node{value: 20, height: 1}

				return &AVLTree{
					root: root,
				}
			},
			expected: &node{
				value:  20,
				height: 2,
				left:   &node{value: 10, height: 1},
				right:  &node{value: 30, height: 1},
			},
		},
		"balances left-heavy tree with LL rotation": {
			setup: func() *AVLTree {
				root := &node{value: 30, height: 3}
				root.left = &node{value: 20, height: 2}
				root.left.left = &node{value: 10, height: 1}

				return &AVLTree{
					root: root,
				}
			},
			expected: &node{
				value:  20,
				height: 2,
				left:   &node{value: 10, height: 1},
				right:  &node{value: 30, height: 1},
			},
		},
		"balances left-heavy tree with LR rotation": {
			setup: func() *AVLTree {
				root := &node{value: 30, height: 3}
				root.left = &node{value: 10, height: 2}
				root.left.right = &node{value: 20, height: 1}

				return &AVLTree{
					root: root,
				}
			},
			expected: &node{
				value:  20,
				height: 2,
				left:   &node{value: 10, height: 1},
				right:  &node{value: 30, height: 1},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := tc.setup()
			balanced := balance(tree.root)

			assert.Equal(t, tc.expected, balanced)
		})
	}
}

func TestInsertAVL(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *AVLTree
		input    int
		expected *AVLTree
	}{
		"inserts to empty tree": {
			setup: func() *AVLTree {
				return NewAVLTree()
			},
			input: 10,
			expected: &AVLTree{
				root: &node{value: 10, height: 1},
			},
		},
		"inserts smaller element": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				return tree
			},
			input: 5,
			expected: &AVLTree{
				root: &node{
					value:  10,
					height: 2,
					left:   &node{value: 5, height: 1},
				},
			},
		},
		"inserts smaller element with rotation": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			},
			input: 2,
			expected: &AVLTree{
				root: &node{
					value:  5,
					height: 2,
					left: &node{
						value:  2,
						height: 1,
					},
					right: &node{
						value:  10,
						height: 1,
					},
				},
			},
		},
		"inserts larger element": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				return tree
			},
			input: 15,
			expected: &AVLTree{
				root: &node{
					value:  10,
					height: 2,
					right:  &node{value: 15, height: 1},
				},
			},
		},
		"inserts larger element with rotation": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(15)
				return tree
			},
			input: 20,
			expected: &AVLTree{
				root: &node{
					value:  15,
					height: 2,
					left: &node{
						value:  10,
						height: 1,
					},
					right: &node{
						value:  20,
						height: 1,
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := tc.setup()
			tree.Insert(tc.input)
			assert.Equal(t, tc.expected, tree)
		})
	}
}

func TestDeleteAVL(t *testing.T) {
	testCases := map[string]struct {
		setup    func() *AVLTree
		input    int
		deleted  bool
		expected *AVLTree
	}{
		"deletes from empty tree": {
			setup: func() *AVLTree {
				return NewAVLTree()
			},
			input:    10,
			deleted:  false,
			expected: &AVLTree{root: nil},
		},
		"deletes root node": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				return tree
			},
			input:   10,
			deleted: true,
			expected: &AVLTree{
				root: nil,
			},
		},
		"deletes leaf node": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			},
			input:   5,
			deleted: true,
			expected: &AVLTree{
				root: &node{
					value:  10,
					height: 1,
				},
			},
		},
		"deletes node with left child only": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			},
			input:   10,
			deleted: true,
			expected: &AVLTree{
				root: &node{
					value:  5,
					height: 1,
				},
			},
		},
		"deletes node with right child only": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(15)
				return tree
			},
			input:   10,
			deleted: true,
			expected: &AVLTree{
				root: &node{
					value:  15,
					height: 1,
				},
			},
		},
		"deletes node with two children and deep left subtree": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(30)
				tree.Insert(20)
				tree.Insert(40)
				tree.Insert(10)
				tree.Insert(25)
				tree.Insert(35)
				tree.Insert(50)
				tree.Insert(5)
				tree.Insert(15)
				tree.Insert(23)
				tree.Insert(27)
				return tree
			},
			input:   30,
			deleted: true,
			expected: &AVLTree{
				root: &node{
					value:  27,
					height: 4,
					left: &node{
						value:  20,
						height: 3,
						left: &node{
							value:  10,
							height: 2,
							left:   &node{value: 5, height: 1},
							right:  &node{value: 15, height: 1},
						},
						right: &node{
							value:  25,
							height: 2,
							left:   &node{value: 23, height: 1},
						},
					},
					right: &node{
						value:  40,
						height: 2,
						left:   &node{value: 35, height: 1},
						right:  &node{value: 50, height: 1},
					},
				},
			},
		},
		"deletes non-existant node": {
			setup: func() *AVLTree {
				tree := NewAVLTree()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			},
			input:   15,
			deleted: false,
			expected: &AVLTree{
				root: &node{
					value:  10,
					height: 2,
					left:   &node{value: 5, height: 1},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tree := tc.setup()
			tree.Delete(tc.input)
			assert.Equal(t, tc.expected, tree)
		})
	}
}

func TestSearchAVL(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	expectFalse := tree.Search(20)
	expectTrue := tree.Search(5)

	assert.False(t, expectFalse)
	assert.True(t, expectTrue)
}

func TestTraverseAVL(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	expected := []int{5, 10, 15}
	assert.Equal(t, expected, tree.Traverse(InOrder))
}
