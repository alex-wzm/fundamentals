package trees

type BinaryTree struct {
	root *Node
}

type Node struct {
	value       int
	left, right *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

// Insert recursively adds a node to the tree while ignoring duplicates
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{value: v}
	}

	t.root.insert(v)
}

func (n *Node) insert(v int) {
	if v < n.value {
		if n.left != nil {
			n.left.insert(v)
		} else {
			n.left = &Node{value: v}
		}
		return
	}

	if v > n.value {
		if n.right != nil {
			n.right.insert(v)
			return
		} else {
			n.right = &Node{value: v}
			return
		}
	}
}

// Search recursively searches for whether the value is in the tree
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Search(v int) bool {
	exists, _ := t.root.search(v)

	return exists
}

func (n *Node) search(v int) (bool, *Node) {
	if v == n.value {
		return true, n
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

	return false, nil
}

// Delete removes the target node and rearranges possible subtrees with the maximum predecessor of the target node
// Time complexity: O(log n) assuming a balanced tree, O(h) in the worst case
// Space complexity: O(log n) assuming a balanced tree, O(h) in the worst case
func (t *BinaryTree) Delete(v int) bool {
	root, deleted := t.root.delete(v)

	t.root = root

	return deleted
}

func (n *Node) delete(v int) (*Node, bool) {
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
		n.left, deleted = n.left.delete(n.value)
	}

	return n, deleted
}

// maxPredecessor returns the maximum value in the left subtree
func (n *Node) maxPredecessor() int {
	maxSoFar := n.left

	for maxSoFar.right != nil {
		maxSoFar = maxSoFar.right
	}

	return maxSoFar.value
}

type TraversalMethod string

const (
	InOrder   TraversalMethod = "InOrder"
	PreOrder  TraversalMethod = "PreOrder"
	PostOrder TraversalMethod = "PostOrder"
)

func (t *BinaryTree) Traverse(method TraversalMethod) []int {
	switch method {
	case InOrder:
		return t.root.traverseInOrder()
	case PreOrder:
		return t.root.traversePreOrder()
	case PostOrder:
		return t.root.traversePostOrder()
	default:
		return t.root.traverseInOrder()
	}
}

func (n *Node) traverseInOrder() []int {
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

func (n *Node) traversePreOrder() []int {
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

func (n *Node) traversePostOrder() []int {
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
