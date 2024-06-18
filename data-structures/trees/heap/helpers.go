package heap

import (
	"errors"
	"math"
)

/**
This package implements a max binary heap.
**/

// parent returns the index of the given element's parent
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the given element's left child
func left(i int) int {
	return 2*i + 1
}

// right returns the index of the given element's right child
func right(i int) int {
	return 2*i + 2
}

// get returns the element at the given index
func (h *Heap) get(i int) int {
	return h.arr[i]
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Size() int {
	return len(h.arr)
}

func (h *Heap) first() (int, error) {
	if h.Size() < 1 {
		return math.MinInt, errors.New("heap is empty")
	}

	return h.get(0), nil
}

// bubbleUp rearranges the element at the given index to it's correct position
func (h *Heap) bubbleUp(i int) {
	for i > 0 && h.get(i) > h.get(parent(i)) {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

// bubbleDown rearranges the element at the given index to it's correct position
func (h *Heap) bubbleDown(i int) {

	// while the heap has at least one more child
	for left(i) < h.Size() {
		// find the index of the largest child
		max := h.maxChild(i)

		// stop bubbling-down when the element is greater than its largest child
		if h.get(i) >= h.get(max) {
			return
		}

		h.swap(i, max)
		i = max
	}

}

// maxChild returns the index of the largest of the element's children
func (h *Heap) maxChild(i int) int {
	max := left(i)
	if right(i) < h.Size() && h.get(right(i)) > h.get(max) {
		max = right(i)
	}
	return max
}
