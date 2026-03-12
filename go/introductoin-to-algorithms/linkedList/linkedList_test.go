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
	last := all[lL.Head.Len()-1]

	if last != valToBeInserted {
		t.Fatalf("The last value in all should %v equal valToBeInserted %v", last, valToBeInserted)
	}

	lRange := 13
	for i := range lRange {
		lL.InsertLast(i + 1)
	}

	allLoop := lL.GetAll()
	lastLoop := allLoop[lL.Head.Len()-1]

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
