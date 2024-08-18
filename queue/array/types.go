package array

import (
	"github.com/vedashruta/go-ds/lists/arraylist"
	"github.com/vedashruta/go-ds/queue"
)

var _ queue.Queue[int] = (*Queue[int])(nil)

type Queue[V comparable] struct {
	arrayList *arraylist.ArrayList[V]
}
