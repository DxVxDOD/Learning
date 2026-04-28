package doublylinkedlist

import (
	"testing"
)

func TestWalkTo(t *testing.T) {
	dl := DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})

	dlNode := dl.walkTo(4)

	if dlNode.item != 5 {
		t.Fatalf("item %v should equal the 4th element", dlNode.item)
	}
}
