package heap

/**
This package implements a max binary heap.
**/

type Heap struct {
	arr []int
}

func New() *Heap {
	return &Heap{}
}

// Insert adds an element of value v to the heap.
// Duplicates are accepted
func (h *Heap) Insert(v int) {
	h.arr = append(h.arr, v)
	h.bubbleUp(h.Size() - 1)
}

// FindMax returns the max element in the heap
func (h *Heap) FindMax() (int, error) {
	return h.first()
}

// ExtractMax removes and returns the max element from the heap
func (h *Heap) ExtractMax() (int, error) {

	max, err := h.first()
	if err != nil {
		return max, err
	}

	// swap the first and last elements
	h.swap(0, h.Size()-1)

	// remove the last element from the array
	h.arr = h.arr[:h.Size()-1]

	h.bubbleDown(0)

	return max, nil
}
