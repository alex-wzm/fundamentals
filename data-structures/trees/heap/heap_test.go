package heap

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	h := New()
	h.Insert(5)
	h.Insert(1)
	h.Insert(1) // intentional duplicate
	h.Insert(2)
	h.Insert(4)
	h.Insert(3)
	h.Insert(2) // intentional duplicate

	assert.Equal(t, 7, h.Size(), "Expected heap size to be 7")
}

func TestFindMax(t *testing.T) {
	h := New()
	max, err := h.FindMax()
	if err != nil {
		assert.Error(t, err, "Expected error, got nil")
	}

	assert.Equal(t, math.MinInt, max, "Expected max element to be -1")

	h.Insert(1)
	h.Insert(2)
	h.Insert(1) // intentional duplicate
	h.Insert(4)
	h.Insert(2) // intentional duplicate
	h.Insert(5)
	h.Insert(3)

	max, err = h.FindMax()
	if err != nil {
		assert.NoError(t, err, "Expected no error")
	}

	assert.Equal(t, 5, max, "Expected max element to be 5")
}

func TestExtractMax(t *testing.T) {
	h := New()
	max, err := h.ExtractMax()
	if err != nil {
		assert.Error(t, err, "Expected error, got nil")
	}

	assert.Equal(t, math.MinInt, max, "Expected max element to be -1")

	h.Insert(-2)
	h.Insert(0)
	h.Insert(5)
	h.Insert(1)
	h.Insert(0) // intentional duplicate
	h.Insert(4)

	max, err = h.ExtractMax()
	if err != nil {
		assert.NoError(t, err, "Expected no error")
	}

	assert.Equal(t, 5, max, "Expected max element to be 5")
	assert.Equal(t, 5, h.Size(), "Expected heap size to be 5")

	max, err = h.ExtractMax()
	if err != nil {
		assert.NoError(t, err, "Expected no error")
	}

	assert.Equal(t, 4, max, "Expected max element to be 4")
	assert.Equal(t, 4, h.Size(), "Expected heap size to be 4")
}
