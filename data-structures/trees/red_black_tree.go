package trees

/** IMPLEMENTATION INCOMPLETE **/

type RedBlackTree struct {
	root *node
}

func NewRedBlackTree() Tree {
	return &RedBlackTree{
		root: nil,
	}
}

// changeColor toggles the node's color
func changeColor(n *node) {
	if n.black {
		n.black = false
	} else {
		n.black = true
	}
}

func (n *node) isRed() bool {
	return !n.black
}

func (n *node) isBlack() bool {
	return n.black
}

func balanceByColor(n *node) *node {
	if n.left == nil && n.right == nil {
		// freshly inserted leaf nodes are red by default before the tree is rebalanced
		return n
	}

	// case 1: black node with red children
	if n.isBlack() && n.left.isRed() && n.right.isRed() {
		changeColor(n)
		changeColor(n.left)
		changeColor(n.right)
		return n
	}

	// case 2: red node with black left child and red right child
	if n.isRed() && (n.left == nil || n.left.isBlack()) && n.right.isRed() {
		return n.rotateLeft(changeColor)
	}

	if n.isRed() && n.left.isRed() && (n.right == nil || n.right.isBlack()) {
		return n.rotateRight(changeColor)
	}

	return n
}

// Insert adds a node of value v to the tree and ignores duplicates
func (t *RedBlackTree) Insert(v int) {
	if t.root == nil {
		t.root = newNode(v)
		changeColor(t.root) // the root node must be black
		return
	}

	t.root = t.root.insert(v, balanceByColor)
}

// Delete removes a node of value v from the tree
func (t *RedBlackTree) Delete(v int) bool {
	if t.root == nil {
		return false
	}

	root, deleted := t.root.delete(v, balanceByColor)

	t.root = root

	return deleted
}

// Search returns whether the value v is in the tree
func (t *RedBlackTree) Search(v int) bool {
	_, exists := t.root.search(v)

	return exists
}

// Traverse returns the tree values in the specified order
func (t *RedBlackTree) Traverse(method TraversalMethod) []int {
	return t.root.traverse(method)
}
