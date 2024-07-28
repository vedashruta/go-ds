package stack

type Stack[V comparable] interface {
	Push(element V)
	Pop() (element V, ok bool)
	Peek() (element V, ok bool)
}
