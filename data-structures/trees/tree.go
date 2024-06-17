package trees

type Tree interface {
	Insert(v int)
	Delete(v int) bool
	Search(v int) bool
	Traverse(TraversalMethod) []int
}
