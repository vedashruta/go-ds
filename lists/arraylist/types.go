package arraylist

import "github.com/vedashruta/go-ds/lists"

var _ lists.List[int] = (*ArrayList[int])(nil)

type ArrayList[V comparable] struct {
	values []V
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)
