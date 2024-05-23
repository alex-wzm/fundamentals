package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue()

	assert.True(q.IsEmpty(), "New queue should be empty")
	assert.Nil(q.Peek(), "Peek should return nil on empty queue")
	assert.Nil(q.Dequeue(), "Dequeue should return nil on empty queue")

	q.Enqueue(1)
	assert.False(q.IsEmpty(), "Queue should not be empty after enqueue")
	assert.Equal(1, q.Size(), "Queue size should be 1 after first enqueue")
	assert.Equal(1, q.Peek(), "Front of queue should be 1")

	q.Enqueue(2)
	assert.Equal(2, q.Size(), "Queue size should be 2 after second enqueue")
	assert.Equal(1, q.Peek(), "Front of queue should still be 1")

	val := q.Dequeue()
	assert.Equal(1, val, "Dequeue should return 1")
	assert.Equal(1, q.Size(), "Queue size should be 1 after dequeue")
	assert.Equal(2, q.Peek(), "Front of queue should be 2")

	q.Dequeue()
	assert.True(q.IsEmpty(), "Queue should be empty after dequeuing last element")
}
