package trees

type node struct {
	value       int
	left, right *node
	height      int
}

func newNode(v int) *node {
	return &node{
		value:  v,
		height: 1,
	}
}

func height(n *node) int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *node) updateHeight() {
	n.height = 1 + max(height(n.left), height(n.right))
}

type callback func(n *node) *node

// noop is a callback that does nothing
func noop(n *node) *node {
	return n
}

type TraversalMethod string

const (
	InOrder   TraversalMethod = "InOrder"
	PreOrder  TraversalMethod = "PreOrder"
	PostOrder TraversalMethod = "PostOrder"
)
