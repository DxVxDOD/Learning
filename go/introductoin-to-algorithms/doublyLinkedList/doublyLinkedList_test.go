package doublylinkedlist_test

import (
	"slices"
	"testing"

	doublylinkedlist "ita/doublyLinkedList"
)

func TestBuild(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})

	all := dl.GetAll()
	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}
	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(all, first) {
		t.Fatalf("all should contain the first element %v", first)
	}

	if !slices.Contains(all, last) {
		t.Fatalf("all should contain the last element %v", last)
	}
}
