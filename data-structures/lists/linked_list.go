package lists

import (
	"fmt"
)

/*
By making linkedList and listNode private, this package relies on the constructor
to ensure the root node is not nil. But there's a guard for that anyway just in case.
*/

type linkedList struct {
	// root node is a "Sentinel" AKA "dummy" node
	root   *Node
	length int
}

type Node struct {
	value any
	next  *Node
}

func NewLinkedList() *linkedList {
	return &linkedList{
		root: &Node{},
	}
}

// Append adds the value `v` to the end of the list
func (l *linkedList) Append(v any) {

	current := l.root

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{value: v}
	l.length++
}

// Insert adds a value before the first instance of the target in the list.
// If no target is found, the value is added to the start of the list.
func (l *linkedList) Insert(v any, target any) {

	current := l.root

	for current.next != nil {
		if current.next.value == target {
			current.next = &Node{
				value: v,
				next:  current.next,
			}
			l.length++
			return
		}

		current = current.next
	}

	// If target is not found, add the value to the start of the list
	l.root.next = &Node{
		value: v,
		next:  l.root.next,
	}
	l.length++
}

// Delete removes the first instance of the target in the list.
func (l *linkedList) Delete(target any) {

	current := l.root

	for current.next != nil {
		if current.next.value == target {
			// Remove the target node by skipping it
			current.next = current.next.next
			l.length--
			return
		}

		current = current.next
	}

}

func (l *linkedList) Get(index int) (any, error) {

	if index > l.length-1 {
		return nil, fmt.Errorf("index out of range [%d] with length %d", index, l.length)
	}

	current := l.root
	i := 0

	for current.next != nil {
		current = current.next

		if i == index {
			return current.value, nil
		}

		i++
	}

	return nil, nil
}

// Traverse returns a slice of all the values in the list
func (l *linkedList) Traverse() []any {

	values := []any{}

	current := l.root

	for current.next != nil {
		current = current.next
		values = append(values, current.value)
	}

	return values
}

// Length returns the number of nodes in the list
func (l *linkedList) Length() int {
	return l.length
}
