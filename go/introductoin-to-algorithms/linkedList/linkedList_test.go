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
