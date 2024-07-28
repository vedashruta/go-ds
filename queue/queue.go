package queue

type Queue[V comparable] interface {
	Enqueue(data V)
	Dequeue() (data V, ok bool)
	Peek() (data V, ok bool)
}
