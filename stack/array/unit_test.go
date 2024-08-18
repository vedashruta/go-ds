package arraystack

import "testing"

func TestStack(t *testing.T) {
	stack := New[int]()
	for i := range 1000 {
		stack.Push(i)
	}
	iterations := 1000
	for i := range iterations {
		data, ok := stack.Peek()
		if !ok {
			t.Fail()
		}
		if data != iterations-(i+1) {
			t.Fail()
		}
		data, ok = stack.Pop()
		if !ok {
			t.Fail()
		}
		if data != iterations-(i+1) {
			t.Fail()
		}
	}
	if stack.arrayList.Size() != 0 {
		t.Fail()
	}

}
