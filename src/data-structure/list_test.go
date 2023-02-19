package list_test

import (
	"container/list"
	"testing"
)

func TestGenerateList(t *testing.T) {
	members := list.New()
	for i, _ := range [10]int{} {
		members.PushBack(i)
	}

	index := 0
	for e := members.Front(); e != nil; e = e.Next() {
		if e.Value != index {
			t.Error("the value in the list doesn't equal to the value")
		}

		index++
	}
}
