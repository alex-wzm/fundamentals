package trees

type BinaryTree struct {
	root *node
}

func NewBinaryTree() Tree {
	return &BinaryTree{
		root: nil,
	}
}

// Insert recursively adds a node to the tree and ignores duplicates
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Insert(v int) {
	if t.root == nil {
		t.root = newNode(v)
	}

	t.root.insert(v, noop)
}

// insert does the binary search tree insertion algorithm
// and takes a callback function to enable AVL tree balancing
func (n *node) insert(v int, callback callback) *node {
	if v < n.value {
		if n.left != nil {
			n.left.insert(v, callback)
		} else {
			n.left = newNode(v)
		}
	}

	if v > n.value {
		if n.right != nil {
			n.right.insert(v, callback)
		} else {
			n.right = newNode(v)
		}
	}

	return callback(n)
}

// Search recursively searches for whether the value is in the tree
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Search(v int) bool {
	_, exists := t.root.search(v)

	return exists
}

func (n *node) search(v int) (*node, bool) {
	if v == n.value {
		return n, true
	}

	if v < n.value {
		if n.left != nil {
			return n.left.search(v)
		}
	}

	if v > n.value {
		if n.right != nil {
			return n.right.search(v)
		}
	}

	return nil, false
}

// Delete removes the target node and recursively rearranges possible subtrees with the maximum predecessor of the target node
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Delete(v int) bool {
	root, deleted := t.root.delete(v, noop)

	t.root = root

	return deleted
}

// delete does the binary search tree deletion algorithm
// and takes a callback function to enable AVL tree balancing
func (n *node) delete(v int, callback callback) (*node, bool) {
	if n == nil {
		return nil, false
	}

	var deleted bool

	if v < n.value {
		n.left, deleted = n.left.delete(v, callback)
	}

	if v > n.value {
		n.right, deleted = n.right.delete(v, callback)
	}

	if v == n.value {
		// no children
		if n.left == nil && n.right == nil {
			return nil, true
		}

		// left child only
		if n.left != nil && n.right == nil {
			return n.left, true
		}

		// right child only
		if n.left == nil && n.right != nil {
			return n.right, true
		}

		// two children
		n.value = n.maxPredecessor()
		n.left, deleted = n.left.delete(n.value, callback)
	}

	return callback(n), deleted
}

// maxPredecessor returns the maximum value in the left subtree
func (n *node) maxPredecessor() int {
	maxSoFar := n.left

	for maxSoFar.right != nil {
		maxSoFar = maxSoFar.right
	}

	return maxSoFar.value
}

// Traverse returns the values of the tree in the specified order
func (t *BinaryTree) Traverse(method TraversalMethod) []int {
	return t.root.traverse(method)
}

// traverse switches between the different traversal methods
func (n *node) traverse(method TraversalMethod) []int {
	switch method {
	case InOrder:
		return n.traverseInOrder()
	case PreOrder:
		return n.traversePreOrder()
	case PostOrder:
		return n.traversePostOrder()
	default:
		return n.traverseInOrder()
	}
}

func (n *node) traverseInOrder() []int {
	var values []int

	if n.left != nil {
		values = append(values, n.left.traverseInOrder()...)
	}

	values = append(values, n.value)

	if n.right != nil {
		values = append(values, n.right.traverseInOrder()...)
	}

	return values
}

func (n *node) traversePreOrder() []int {
	var values []int

	values = append(values, n.value)

	if n.left != nil {
		values = append(values, n.left.traversePreOrder()...)
	}

	if n.right != nil {
		values = append(values, n.right.traversePreOrder()...)
	}

	return values
}

func (n *node) traversePostOrder() []int {
	var values []int

	if n.left != nil {
		values = append(values, n.left.traversePostOrder()...)
	}

	if n.right != nil {
		values = append(values, n.right.traversePostOrder()...)
	}

	values = append(values, n.value)

	return values
}
