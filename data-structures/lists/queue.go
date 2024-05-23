package lists

type Queue struct {
	root   *Node
	tail   *Node
	length int
}

func NewQueue() *Queue {
	sentinel := &Node{}

	return &Queue{
		root: sentinel,
		tail: sentinel,
	}
}

func (q *Queue) Enqueue(v any) {
	newNodePointer := &Node{value: v}
	q.tail.next = newNodePointer
	q.tail = newNodePointer
	q.length++
}

func (q *Queue) Dequeue() any {
	if q.length > 0 {
		node := q.root.next
		if node.next != nil {
			q.root.next = node.next
		}
		q.length--

		// reset tail when dequeuing last node
		if q.length == 0 {
			q.tail = q.root
		}
		return node.value
	}

	return nil
}

func (q *Queue) Peek() any {
	if q.length > 0 {
		return q.root.next.value
	}

	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Size() int {
	return q.length
}
