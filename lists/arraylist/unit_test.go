package arraylist

import (
	"testing"
)

func TestAddRemove(t *testing.T) {
	q := New[int]()
	for i := range 1000 {
		q.Add(i)
	}
	if len(q.values) != 1000 {
		t.Fail()
	}
	for i := range 1000 {
		val, ok := q.Get(i)
		if !ok {
			t.Fail()
		}
		if val != i {
			t.Fail()
		}
	}
	for {
		if len(q.values) > 0 {
			q.Remove(0)
		} else {
			break
		}
	}
	if len(q.values) != 0 {
		t.Fail()
	}
}

func TestGetSet(t *testing.T) {
	q := New[int]()
	for i := range 1000 {
		q.Set(i, i)
	}
	if len(q.values) != 1000 {
		t.Fail()
	}
	for i := range 1000 {
		value, ok := q.Get(i)
		if !ok {
			t.Fail()
		}
		if value != i {
			t.Fail()
		}
	}
}

func TestIndexOf(t *testing.T) {
	q := New[int]()
	for i := range 1000 {
		q.Add(i)
	}
	if len(q.values) != 1000 {
		t.Fail()
	}
	for i := range 1000 {
		index := q.IndexOf(i)
		if index != i {
			t.Fail()
		}
	}
}

func TestClear(t *testing.T) {
	q := New[int]()
	for i := range 1000 {
		q.Add(i)
	}
	q.Clear()
	if len(q.values) != 0 {
		t.Fail()
	}
	_, ok := q.Get(0)
	if ok {
		t.Fail()
	}
}
