package linkedlist_test

import (
	"slices"
	"testing"

	linkedlist "ita/linkedList"
)

func TestBuildAndGetAll(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1, 3, 4, 5, 2, 6, 8, 5, 6, 9, 11})

	all := lL.GetAll()

	randItem := 8
	if !slices.Contains(all, randItem) {
		t.Fatalf("did not build the slice correctly, missing %v ", randItem)
	}

	firstItem := 1
	if !slices.Contains(all, firstItem) {
		t.Fatalf("did not build the slice correctly, missing %v ", firstItem)
	}

	lastItem := 11
	if !slices.Contains(all, lastItem) {
		t.Fatalf("did not build the slice correctly, missing %v ", lastItem)
	}
}

func TestInsertLast(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1})

	valToBeInserted := 10
	lL.InsertLast(valToBeInserted)

	all := lL.GetAll()
	last := all[lL.Len()-1]

	if last != valToBeInserted {
		t.Fatalf("The last value in all should %v equal valToBeInserted %v", last, valToBeInserted)
	}

	lRange := 13
	for i := range lRange {
		lL.InsertLast(i + 1)
	}

	allLoop := lL.GetAll()
	lastLoop := allLoop[lL.Len()-1]

	if lastLoop != lRange {
		t.Fatalf("The last value %v in allLoop should equal lRange %v", lastLoop, lRange)
	}
}

func TestGetLast(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1})

	lRange := 13
	for i := range lRange {
		lL.InsertLast(i + 1)
	}

	lastNode := lL.GetLast()
	lastValue := lastNode.Item

	if lRange != lastValue {
		t.Fatalf("The last value %v should equal lRange %v", lastValue, lRange)
	}
}

func TestDeleteLast(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1})

	lRange := 13
	for i := range lRange {
		lL.InsertLast(i + 1)
	}

	lastNode := lL.GetLast()
	lastValue := lastNode.Item

	if lRange != lastValue {
		t.Fatalf("The last value %v should equal lRange %v", lastValue, lRange)
	}

	deletedNode, err := lL.DeleteLast()
	if err != nil {
		t.Fatal(err)
	}

	newLastNode := lL.GetLast()

	newLastNodeItem := newLastNode.Item
	deletedNodeItem := deletedNode.Item

	if newLastNodeItem == deletedNodeItem {
		t.Fatalf("The new las item %v should equal the deleted node item %v", newLastNodeItem, deletedNodeItem)
	}

	if newLastNodeItem+1 != deletedNodeItem {
		t.Fatalf("The new las item %v should be smaller by (per tests logic) the deleted node item %v", newLastNodeItem, deletedNodeItem)
	}
}

func TestInsertFirst(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1, 2, 3, 4})

	valToBeInserted := 10
	newHead := lL.InsertFirst(valToBeInserted)
	newHeadItem := newHead.Item

	all := lL.GetAll()
	first := all[0]

	if first != newHeadItem {
		t.Fatalf("first %v should equal newHeadItem %v", first, newHeadItem)
	}

	getFirst, err := lL.GetFirst()
	if err != nil {
		t.Fatal(err)
	}
	getFirstItem := getFirst.Item

	if getFirstItem != newHeadItem {
		t.Fatalf("first %v should equal newHeadItem %v", first, newHeadItem)
	}
}

func TestGetAt(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})

	linkAtIdx, err := lL.GetAt(3)
	if err != nil {
		t.Fatal(err)
	}

	linkItemAtIdx := linkAtIdx.Item
	expectedValue := 4
	if linkItemAtIdx != expectedValue {
		t.Fatalf("valAtIdx %v does not match the expected value %v", linkItemAtIdx, expectedValue)
	}

	_, err = lL.GetAt(10)
	if err == nil {
		t.Fatal("this Should be index out of bounds")
	}
}

func TestInsertAt(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1, 2})

	for i := range 5 {
		_, err := lL.InsertAt(i+1, i)
		if err != nil {
			t.Fatal(err)
		}
	}

	all := lL.GetAll()

	if !slices.Contains(all, 5) {
		t.Fatal("should contain 5")
	}

	_, err := lL.InsertAt(11, 4)
	if err != nil {
		t.Fatal(err)
	}

	valAt, err := lL.GetAt(4)
	if err != nil {
		t.Fatal(err)
	}
	if valAt.Item != 11 {
		t.Fatalf("valAt %v should equal 11", valAt.Item)
	}
}

func TestDeleteAt(t *testing.T) {
	starter := linkedlist.Linkedlist[int]{}
	lL := starter.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})
	starterLen := lL.Len()

	_, err := lL.DeleteAt(3)
	if err != nil {
		t.Fatal(err)
	}

	all := lL.GetAll()

	if slices.Contains(all, 3) {
		t.Fatal("The linked list should not contain 3")
	}

	afterLen := lL.Len()

	if starterLen == afterLen {
		t.Fatalf("starterLen %v should not equal afterLen %v", starterLen, afterLen)
	}

	diff := starterLen - afterLen
	if diff != 1 {
		t.Fatalf("The diff %v should be 1", diff)
	}
}
