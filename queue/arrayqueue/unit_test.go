package arrayqueue

import "testing"

func TestQueue(t *testing.T) {
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
