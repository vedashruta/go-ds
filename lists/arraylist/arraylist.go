package arraylist

import (
	"slices"

	"github.com/vedashruta/go-ds/lists"
)

var _ lists.List[int] = (*ArrayList[int])(nil)

type ArrayList[V comparable] struct {
	values []V
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

func New[V comparable]() (arrayList *ArrayList[V]) {
	arrayList = &ArrayList[V]{}
	return
}

func (arrayList *ArrayList[V]) resize(length int, capacity int) {
	newValues := make([]V, length, capacity)
	copy(newValues, arrayList.values)
	arrayList.values = newValues
}

func (arrayList *ArrayList[V]) withinRange(index int) (ok bool) {
	if index >= 0 && index < len(arrayList.values) {
		ok = true
		return
	}
	return
}

func (arraylist *ArrayList[V]) growBy(n int) {
	currentCapacity := cap(arraylist.values)
	newLength := len(arraylist.values) + n
	if newLength >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		arraylist.resize(newLength, newCapacity)
	} else {
		arraylist.values = arraylist.values[:newLength]
	}
}

func (arraylist *ArrayList[V]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(arraylist.values)
	if len(arraylist.values) <= int(float32(currentCapacity)*shrinkFactor) {
		arraylist.resize(len(arraylist.values), len(arraylist.values))
	}
}

func (arraylist *ArrayList[V]) Add(values ...V) {
	length := len(arraylist.values)
	arraylist.growBy(len(values))
	for i := range values {
		arraylist.values[length+i] = values[i]
	}
}

func (arraylist *ArrayList[V]) Remove(index int) {
	if arraylist.withinRange(index) {
		arraylist.values = slices.Delete(arraylist.values, index, index+1)
		arraylist.shrink()
	}
}

func (arraylist *ArrayList[V]) Get(index int) (value V, ok bool) {
	if arraylist.withinRange(index) {
		value = arraylist.values[index]
		ok = true
	}
	return
}

func (arraylist *ArrayList[V]) Set(value V, index int) {
	if arraylist.withinRange(index) {
		if index == len(arraylist.values) {
			arraylist.Add(value)
			return
		}
	}
	arraylist.values[index] = value
}

func (arraylist *ArrayList[V]) IndexOf(value V) (index int) {
	index = slices.Index(arraylist.values, value)
	return
}

// arraylist.Clear() will clear all the elelemts in the given arraylist,if the capacity of the slice was n,then
// the capacity of the slice cap(arraylist.values) remains n
func (arraylist *ArrayList[V]) Clear() {
	clear(arraylist.values[:cap(arraylist.values)])
	arraylist.values = arraylist.values[:0]
}

func (arraylist *ArrayList[V]) Size() (size int) {
	size = len(arraylist.values)
	return
}
