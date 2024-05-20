package lists

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {

	testCases := map[string]struct {
		setup    func() *linkedList
		input    any
		expected *linkedList
	}{
		"empty": {
			setup:    NewLinkedList,
			input:    "hello",
			expected: &linkedList{root: &Node{next: &Node{value: "hello"}}, length: 1},
		},
		"basic": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				return list
			},
			input:    "second",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second"}}}, length: 2},
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			list := c.setup()

			list.Append(c.input)

			assert.Equal(t, c.expected, list)

		})
	}
}

func TestInsert(t *testing.T) {

	testCases := map[string]struct {
		setup    func() *linkedList
		input    any
		target   any
		expected *linkedList
	}{
		"insert in the beginning": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("second")
				list.Append("third")
				return list
			},
			input:    "first",
			target:   "second",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second", next: &Node{value: "third"}}}}, length: 3},
		},
		"insert in the middle": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("third")
				return list
			},
			input:    "second",
			target:   "third",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second", next: &Node{value: "third"}}}}, length: 3},
		},
		"insert into empty list": {
			setup:    NewLinkedList,
			input:    "first",
			target:   "second",
			expected: &linkedList{root: &Node{next: &Node{value: "first"}}, length: 1},
		},
		"inserts before first instance of target": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("second")
				list.Append("second")
				return list
			},
			input:    "first",
			target:   "second",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second", next: &Node{value: "second"}}}}, length: 3},
		},
		"target not found": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("second")
				list.Append("fourth")
				return list
			},
			input:    "first",
			target:   "third",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second", next: &Node{value: "fourth"}}}}, length: 3},
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			list := c.setup()

			list.Insert(c.input, c.target)

			assert.Equal(t, c.expected, list)

		})
	}
}

func TestDelete(t *testing.T) {

	testCases := map[string]struct {
		setup    func() *linkedList
		target   any
		expected *linkedList
	}{
		"delete from empty list": {
			setup:    NewLinkedList,
			target:   "hello",
			expected: &linkedList{root: &Node{}, length: 0},
		},
		"delete from single-node list": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("hello")
				return list
			},
			target:   "hello",
			expected: &linkedList{root: &Node{}, length: 0},
		},
		"delete from middle of list": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				list.Append("third")
				return list
			},
			target:   "second",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "third"}}}, length: 2},
		},
		"delete from end of list": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				list.Append("third")
				return list
			},
			target:   "third",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second"}}}, length: 2},
		},
		"delete non-existing target": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				return list
			},
			target:   "third",
			expected: &linkedList{root: &Node{next: &Node{value: "first", next: &Node{value: "second"}}}, length: 2},
		},
	}
	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			list := c.setup()

			list.Delete(c.target)

			assert.Equal(t, c.expected, list)

		})
	}
}

func TestGet(t *testing.T) {

	testCases := map[string]struct {
		setup    func() *linkedList
		index    int
		expected any
		err      error
	}{
		"empty": {
			setup:    NewLinkedList,
			index:    0,
			expected: nil,
			err:      errors.New("index out of range"),
		},
		"basic": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				list.Append("third")
				return list
			},
			index:    1,
			expected: "second",
			err:      nil,
		},
		"index out of range": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				list.Append("third")
				return list
			},
			index:    3,
			expected: nil,
			err:      errors.New("index out of range"),
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			list := c.setup()

			result, err := list.Get(c.index)
			if c.err != nil {
				assert.ErrorContains(t, err, c.err.Error())
			} else {
				assert.NoError(t, c.err)
				assert.Equal(t, c.expected, result)
			}
		})
	}
}

func TestTraverse(t *testing.T) {

	testCases := map[string]struct {
		setup    func() *linkedList
		expected []any
	}{
		"empty": {
			setup:    NewLinkedList,
			expected: []any{},
		},
		"basic": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append("first")
				list.Append("second")
				list.Append("third")
				return list
			},
			expected: []any{"first", "second", "third"},
		},
		"mixed types": {
			setup: func() *linkedList {
				list := NewLinkedList()
				list.Append(1)
				list.Append("second")
				list.Append(3)
				return list
			},
			expected: []any{1, "second", 3},
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			list := c.setup()

			result := list.Traverse()

			assert.Equal(t, c.expected, result)

		})
	}

}

func TestLength(t *testing.T) {
	list := NewLinkedList()

	assert.Equal(t, 0, list.Length())

	list.Append("first")
	list.Append("second")
	list.Append("third")

	assert.Equal(t, 3, list.Length())
}
