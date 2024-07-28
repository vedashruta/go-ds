package arrayqueue

import (
	"github.com/vedashruta/go-ds/lists/arraylist"
	"github.com/vedashruta/go-ds/queue"
)

var _ queue.Queue[int] = (*Queue[int])(nil)

type Queue[V comparable] struct {
	arrayList *arraylist.ArrayList[V]
}

func New[V comparable]() (q *Queue[V]) {
	q = &Queue[V]{arrayList: arraylist.New[V]()}
	return
}

func (q *Queue[V]) Enqueue(value V) {
	q.arrayList.Add(value)
}

func (q *Queue[V]) Dequeue() (value V, ok bool) {
	value, ok = q.arrayList.Get(0)
	if ok {
		q.arrayList.Remove(0)
	}
	return
}

func (q *Queue[V]) Peek() (value V, ok bool) {
	value, ok = q.arrayList.Get(0)
	if ok {
		return
	}
	return
}
