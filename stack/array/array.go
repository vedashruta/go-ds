package arraystack

import (
	"github.com/vedashruta/go-ds/lists/arraylist"
	"github.com/vedashruta/go-ds/stack"
)

var _ stack.Stack[int] = (*Stack[int])(nil)

type Stack[V comparable] struct {
	arrayList *arraylist.ArrayList[V]
}

// Creates a new arrayList of type Stack
func New[V comparable]() (stack *Stack[V]) {
	stack = &Stack[V]{arrayList: arraylist.New[V]()}
	return
}

// Adds the element to the end of the arrayList
// i.e in an arrayList of size n,element will be inserted to n+1th position
func (s *Stack[V]) Push(element V) {
	s.arrayList.Add(element)
}

// Removes the last element that was inserted
// i.e in an arrayList of size n,element will be deleted from n-1th position
func (s *Stack[V]) Pop() (element V, ok bool) {
	element, ok = s.arrayList.Get(s.arrayList.Size() - 1)
	if ok {
		s.arrayList.Remove(s.arrayList.Size() - 1)
		return
	}
	return
}

// Returns the last element that was inserted
// i.e in an arrayList of size n,element from n-1th position will be returned
func (s *Stack[V]) Peek() (element V, ok bool) {
	element, ok = s.arrayList.Get(s.arrayList.Size() - 1)
	if ok {
		return
	}
	return
}
