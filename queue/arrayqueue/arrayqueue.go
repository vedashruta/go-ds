package arrayqueue

import (
	"github.com/vedashruta/go-ds/lists/arraylist"
	"github.com/vedashruta/go-ds/queue"
)

var _ queue.Queue[int] = (*Queue[int])(nil)

type Queue[V comparable] struct {
	arrayList *arraylist.ArrayList[V]
}

// Creates a new arrayList of type Queue
func New[V comparable]() (q *Queue[V]) {
	q = &Queue[V]{arrayList: arraylist.New[V]()}
	return
}

// Adds the element to the end of the arrayList
// i.e in an arrayList of size n,element will be inserted to n+1th position
func (q *Queue[V]) Enqueue(value V) {
	q.arrayList.Add(value)
}

// Deleted the element that was inserted first
// i.e in an arrayList of size n,element will be deleted from ith position where i<=n
func (q *Queue[V]) Dequeue() (value V, ok bool) {
	value, ok = q.arrayList.Get(0)
	if ok {
		q.arrayList.Remove(0)
	}
	return
}

// Returns the first element that was inserted
// i.e in an arrayList of size n,element from ith position will be returned where i<=n
func (q *Queue[V]) Peek() (value V, ok bool) {
	value, ok = q.arrayList.Get(0)
	if ok {
		return
	}
	return
}
