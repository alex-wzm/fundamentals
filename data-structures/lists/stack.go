package lists

// Stack is a linked list following the LIFO principle.
type Stack linkedList

func NewStack() *Stack {
	return &Stack{
		root: &Node{},
	}
}

func (s *Stack) Push(v any) {
	s.root.next = &Node{
		value: v,
		next:  s.root.next,
	}

	s.length++
}

func (s *Stack) Pop() any {

	if s.root.next != nil {

		result := s.root.next.value
		s.root.next = s.root.next.next

		s.length--

		return result
	}

	return nil
}

func (s *Stack) Peek() any {

	if s.root.next != nil {
		return s.root.next.value
	}

	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.root.next == nil
}

func (s *Stack) Size() int {
	return s.length
}
