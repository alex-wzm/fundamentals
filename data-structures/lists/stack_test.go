package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()

	assert.True(t, s.IsEmpty(), "New stack should be empty")

	val := s.Peek()
	assert.Nil(t, val, "Peek should return nil on empty stack")

	val = s.Pop()
	assert.Nil(t, val, "Pop should return nil on empty stack")

	s.Push(1)
	assert.False(t, s.IsEmpty(), "Stack should not be empty after push")
	assert.Equal(t, 1, s.Size(), "Stack size should be 1 after first push")
	assert.Equal(t, 1, s.Peek(), "Top of stack should be 1")

	s.Push(2)
	assert.Equal(t, 2, s.Size(), "Stack size should be 2 after second push")
	assert.Equal(t, 2, s.Peek(), "Top of stack should be 2")

	val = s.Pop()
	assert.Equal(t, 2, val, "Pop should return 2")
	assert.Equal(t, 1, s.Size(), "Stack size should be 1 after pop")
	assert.Equal(t, 1, s.Peek(), "Top of stack should be 1")

	s.Pop()
	assert.True(t, s.IsEmpty(), "Stack should be empty after popping last element")
}
