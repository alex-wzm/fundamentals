package trees

type AVLTree struct {
	root *node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func (n *node) balanceFactor() int {
	return height(n.left) - height(n.right)
}

func (n *node) isLeftHeavy() bool {
	return n.balanceFactor() > 0
}

func (n *node) isRightHeavy() bool {
	return n.balanceFactor() < 0
}

func (n *node) requiresLeftRotation() bool {
	return n.balanceFactor() > 1
}

func (n *node) requiresRightRotation() bool {
	return n.balanceFactor() < -1
}

// balance modifies the tree to satisfy the AVL balance property
// using the four possible rotations: LL, LR, RL, RR
// Time complexity: O(1)
func balance(n *node) *node {
	if n.requiresLeftRotation() {
		if n.left.isRightHeavy() {
			// LR case - rotate left child left before main rotation
			n.left = n.left.rotateLeft()
		}

		// LL case - rotate right
		return n.rotateRight()
	}

	if n.requiresRightRotation() {
		if n.right.isLeftHeavy() {
			// RL case - rotate right child right before main rotation
			n.right = n.right.rotateRight()
		}
		// RR - rotate left
		return n.rotateLeft()
	}

	return n
}

// rotateLeft is performed when a node is right-heavy
func (x *node) rotateLeft() *node {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.updateHeight()
	y.updateHeight()

	return y
}

// rotateRight is performed when a node is left-heavy
func (y *node) rotateRight() *node {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.updateHeight()
	x.updateHeight()

	return x
}

// Insert adds a node of value v to the tree and ignores duplicates
func (t *AVLTree) Insert(v int) {
	if t.root == nil {
		t.root = newNode(v)

		return
	}

	t.root = t.root.insert(v, balance)
}

// Delete removes a node of value v from the tree
func (t *AVLTree) Delete(v int) bool {
	if t.root == nil {
		return false
	}

	root, deleted := t.root.delete(v, balance)

	t.root = root

	return deleted
}

// Search returns whether the value v is in the tree
func (t *AVLTree) Search(v int) bool {
	_, exists := t.root.search(v)

	return exists
}

// Traverse returns the tree values in the specified order
func (t *AVLTree) Traverse(method TraversalMethod) []int {
	return t.root.traverse(method)
}
