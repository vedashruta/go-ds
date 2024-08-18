package array

import (
	"fmt"
	"testing"
)

func TestIntegerQueue(t *testing.T) {
	queue := New[int]()
	iterations := 1000
	for i := range iterations {
		queue.Enqueue(i)
	}
	for i := range iterations {
		data, ok := queue.Peek()
		if !ok {
			t.Fail()
		}
		if data != i {
			t.Fail()
		}
		data, ok = queue.Dequeue()
		if !ok {
			t.Fail()
		}
		if data != i {
			t.Fail()
		}
	}
	if queue.arrayList.Size() != 0 {
		t.Fail()
	}
}

func TestStringQueue(t *testing.T) {
	queue := New[string]()
	iterations := 1000
	for i := range iterations {
		element := fmt.Sprint("%1[s]", i)
		queue.Enqueue(element)
	}
	for i := range iterations {
		data := fmt.Sprint("%1[s]", i)
		element, ok := queue.Peek()
		if !ok {
			t.Fail()
		}
		if element != data {
			t.Fail()
		}
		element, ok = queue.Dequeue()
		if !ok {
			t.Fail()
		}
		if element != data {
			t.Fail()
		}
	}
	if queue.arrayList.Size() != 0 {
		t.Fail()
	}
}
