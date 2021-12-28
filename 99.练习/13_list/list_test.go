package list_test

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	lst := list.New()
	for i := 0; i < 10; i++ {
		lst.PushBack(i)
	}
	for lst.Len() > 0 {
		v := lst.Front()
		t.Log(lst.Remove(v).(int))
	}
}
