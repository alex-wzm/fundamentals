package trees

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	value       int
	left, right *AVLNode
	height      int
	// balanceFactor int // must be either -1, 0, or 1
}

func NewAVLTree() *AVLTree {

	return &AVLTree{
		root: nil,
	}
}

func NewAVLNode(v int) *AVLNode {
	return &AVLNode{
		value:  v,
		height: 1,
	}
}

func height(n *AVLNode) int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *AVLNode) updateHeight() {
	n.height = 1 + max(height(n.left), height(n.right))
}

func (n *AVLNode) balanceFactor() int {
	return height(n.left) - height(n.right)
}

func (n *AVLNode) isLeftHeavy() bool {
	return n.balanceFactor() > 0
}

func (n *AVLNode) isRightHeavy() bool {
	return n.balanceFactor() < 0
}

func (n *AVLNode) requiresLeftRotation() bool {
	return n.balanceFactor() > 1
}

func (n *AVLNode) requiresRightRotation() bool {
	return n.balanceFactor() < -1
}

func (n *AVLNode) balance() *AVLNode {
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
func (x *AVLNode) rotateLeft() *AVLNode {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.updateHeight()
	y.updateHeight()

	return y
}

// rotateRight is performed when a node is left-heavy
func (y *AVLNode) rotateRight() *AVLNode {
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
		t.root = NewAVLNode(v)

		return
	}

	t.root = t.root.insert(v)
}

// insert rebalances the tree after insertion
func (n *AVLNode) insert(v int) *AVLNode {

	if v < n.value {
		if n.left != nil {
			n.left = n.left.insert(v)
		} else {
			n.left = NewAVLNode(v)
		}
	}

	if v > n.value {
		if n.right != nil {
			n.right = n.right.insert(v)
		} else {
			n.right = NewAVLNode(v)
		}
	}

	n.updateHeight()
	return n.balance()
}

func (t *AVLTree) Delete(v int) bool {
	if t.root == nil {
		return false
	}

	root, deleted := t.root.delete(v)

	t.root = root

	return deleted
}

func (n *AVLNode) delete(v int) (*AVLNode, bool) {
	if n == nil {
		return nil, false
	}

	var deleted bool

	if v < n.value {
		n.left, deleted = n.left.delete(v)
	}

	if v > n.value {
		n.right, deleted = n.right.delete(v)
	}

	if v == n.value {
		// no children
		if n.left == nil && n.right == nil {
			return nil, true
		}

		// left only
		if n.left != nil && n.right == nil {
			return n.left, true
		}

		// right only
		if n.left == nil && n.right != nil {
			return n.right, true
		}

		// two children
		n.value = n.maxPredecessor()
		n.left, deleted = n.left.delete(n.value)
	}

	n.updateHeight()

	return n.balance(), deleted

}

func (n *AVLNode) maxPredecessor() int {
	maxSoFar := n.left

	for maxSoFar.right != nil {
		maxSoFar = maxSoFar.right
	}

	return maxSoFar.value
}
