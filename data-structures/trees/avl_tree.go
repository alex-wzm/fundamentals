package trees

type AVLTree struct {
	root *node
}

func NewAVLTree() Tree {
	return &AVLTree{
		root: nil,
	}
}

func height(n *node) int {
	if n == nil {
		return 0
	}

	return n.height
}

func updateHeight(n *node) {
	n.height = 1 + max(height(n.left), height(n.right))
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

// balanceByHeight modifies the tree to satisfy the AVL balance property
// using the four possible rotations: LL, LR, RL, RR
// Time complexity: O(1)
func balanceByHeight(n *node) *node {
	updateHeight(n)

	if n.requiresLeftRotation() {
		if n.left.isRightHeavy() {
			// LR case - rotate left child left before main rotation
			n.left = n.left.rotateLeft(updateHeight)
		}

		// LL case - rotate right
		return n.rotateRight(updateHeight)
	}

	if n.requiresRightRotation() {
		if n.right.isLeftHeavy() {
			// RL case - rotate right child right before main rotation
			n.right = n.right.rotateRight(updateHeight)
		}
		// RR - rotate left
		return n.rotateLeft(updateHeight)
	}

	return n
}

// Insert adds a node of value v to the tree and ignores duplicates
func (t *AVLTree) Insert(v int) {
	if t.root == nil {
		t.root = newNode(v)

		return
	}

	t.root = t.root.insert(v, balanceByHeight)
}

// Delete removes a node of value v from the tree
func (t *AVLTree) Delete(v int) bool {
	if t.root == nil {
		return false
	}

	root, deleted := t.root.delete(v, balanceByHeight)

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
