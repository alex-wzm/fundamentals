package trees

type node struct {
	value       int
	left, right *node
	// height is used for AVL Tree balancing
	height int
	// black is used for Red-Black Tree balancing (nodes are red by default)
	black bool
}

func newNode(v int) *node {
	return &node{
		value:  v,
		height: 1,
		/**
		Idea: Consider having a separate newAVLNode function
		so BST node heights are always 0 (i.e. it is implied that 0 height = not an AVL node)
		**/
	}
}

// callback enables middleware logic for specific use cases
type callback func(n *node) *node

// modifierFn enables abstracted updates to the value of a node in-place
type modifierFn func(n *node)

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

// rotateLeft is performed when a node is right-heavy
func (x *node) rotateLeft(modifierFn modifierFn) *node {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	modifierFn(x)
	modifierFn(y)

	return y
}

// rotateRight is performed when a node is left-heavy
func (y *node) rotateRight(modifierFn modifierFn) *node {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	modifierFn(y)
	modifierFn(x)

	return x
}
