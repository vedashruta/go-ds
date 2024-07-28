package lists

type List[V comparable] interface {
	Add(values ...V)
	Remove(index int)
	Get(index int) (V, bool)
	Set(index int, value V)
}
