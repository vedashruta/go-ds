package arraystack

import (
	"github.com/vedashruta/go-ds/lists/arraylist"
	"github.com/vedashruta/go-ds/stack"
)

var _ stack.Stack[int] = (*Stack[int])(nil)

type Stack[V comparable] struct {
	arrayList *arraylist.ArrayList[V]
}

func New[V comparable]() (stack *Stack[V]) {
	stack = &Stack[V]{arrayList: arraylist.New[V]()}
	return
}

func (s *Stack[V]) Push(element V) {
	s.arrayList.Add(element)
}

func (s *Stack[V]) Pop() (element V, ok bool) {
	element, ok = s.arrayList.Get(s.arrayList.Size() - 1)
	if ok {
		s.arrayList.Remove(s.arrayList.Size() - 1)
		return
	}
	return
}

func (s *Stack[V]) Peek() (element V, ok bool) {
	element, ok = s.arrayList.Get(s.arrayList.Size() - 1)
	if ok {
		return
	}
	return
}
