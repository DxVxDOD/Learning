package dynamicsequence_test

import (
	"testing"

	dynamicsequence "ita/dynamicSequence"
)

func TestBuildAndGetAt(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	starterLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}

	starterCap := dynSeq.Cap()

	last, err := dynSeq.GetAt(starterLen - 1)
	if err != nil {
		t.Fatal(err)
	}
	first, err := dynSeq.GetAt(0)
	if err != nil {
		t.Fatal(err)
	}

	if first != 10 {
		t.Fatalf("should have returned the first value: %v", first)
	}

	if last != 80 {
		t.Fatalf("should have returned the last value: %v", last)
	}

	if starterLen != 8 {
		t.Fatalf("starter length initialised poorly: %v", starterLen)
	}
	if starterCap != starterLen*3 {
		t.Fatalf("starter capacity is not the right size %v", starterCap)
	}
}

func TestUpSizeAndInsertFirst(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	_, err := dynSeq.GetFirst()
	if err == nil {
		t.Fatal("This should have err")
	}

	oldLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}
	oldCap := dynSeq.Cap()

	for i := range dynSeq.Len() {
		_, err := dynSeq.InsertFirst(i + 1)
		if err != nil {
			t.Fatal(err)
		}
	}

	newLen := dynSeq.Len()
	newCap := dynSeq.Cap()

	if oldLen >= newLen {
		t.Fatalf("new length should be larger than the new one. New length: %v", newLen)
	}

	if oldCap*2 != newCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old cap -> %v", newCap, oldCap)
	}

	first, err := dynSeq.GetAt(0)
	if err != nil {
		t.Fatal(err)
	}

	if first != oldLen {
		t.Fatalf("the first value %v does not match what was inserted %v", first, oldLen)
	}

	getFirst, err := dynSeq.GetFirst()
	if err != nil {
		t.Fatal(err)
	}
	if getFirst != oldLen {
		t.Fatalf("the first value %v does not match what was inserted %v getFirst", first, oldLen)
	}

	for i := range dynSeq.Len() {
		_, err := dynSeq.InsertFirst(i + 1)
		if err != nil {
			t.Fatal(err)
		}
	}
	seoondNewCap := dynSeq.Cap()

	if newCap*2 != seoondNewCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old length -> %v", newCap, oldLen)
	}
}

func TestDeleteFirstAndDownSize(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	startLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80, 10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}
	startCap := dynSeq.Cap()

	numOfDeleteFirstLoops := startLen / 2
	for range numOfDeleteFirstLoops {
		toBeDeleted, err := dynSeq.GetAt(0)
		if err != nil {
			t.Fatal(err)
		}
		val, err := dynSeq.DeleteFirst()
		if err != nil {
			t.Fatalf("unexpected error on DeleteFirst: %e", err)
		}
		if toBeDeleted != val {
			t.Fatalf("The deleted value should match the value captured above: toBeDeleted %v <-> %v val", toBeDeleted, val)
		}
	}

	decLen := dynSeq.Len()
	if decLen != startLen-numOfDeleteFirstLoops {
		t.Fatalf("Length should have decreased by %v, curr len %v", numOfDeleteFirstLoops, dynSeq.Len())
	}
	decCap := dynSeq.Cap()
	if startCap/2 != decCap {
		t.Fatalf("starter capacity %v should have decreased by 2 folds, new capacity %v", startCap, decCap)
	}

	numOfDeleteSecondLoops := startLen / 4
	for range numOfDeleteSecondLoops {
		toBeDeleted, err := dynSeq.GetAt(0)
		if err != nil {
			t.Fatal(err)
		}
		val, err := dynSeq.DeleteFirst()
		if err != nil {
			t.Fatalf("unexpected error on DeleteFirst: %e", err)
		}
		if toBeDeleted != val {
			t.Fatalf("The deleted value should match the value captured above: toBeDeleted %v <-> %v val", toBeDeleted, val)
		}
	}

	secondDecLen := dynSeq.Len()
	if secondDecLen != decLen-numOfDeleteSecondLoops {
		t.Fatalf("Length should have decreased by %v, curr len %v", numOfDeleteSecondLoops, dynSeq.Len())
	}
	secondDecCap := dynSeq.Cap()
	if decCap/2 != secondDecCap {
		t.Fatalf("second capacity %v should have decreased by 2 folds, new capacity %v", startCap, decCap/2)
	}
}

func TestInseertLastAndGetLast(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	_, err := dynSeq.GetLast()
	if err == nil {
		t.Fatal("This should have err")
	}

	oldLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}

	oldCap := dynSeq.Cap()

	for i := range dynSeq.Len() {
		_, err := dynSeq.InsertLast(i + 1)
		if err != nil {
			t.Fatal(err)
		}
	}

	newLen := dynSeq.Len()
	newCap := dynSeq.Cap()

	if oldLen >= newLen {
		t.Fatalf("new length should be larger than the new one. New length: %v oldLen: %v", newLen, oldLen)
	}

	if oldCap*2 != newCap {
		t.Fatalf("new capacity should be 2 times the old cap: new capacity -> %v, old cap -> %v", newCap, oldCap)
	}

	last, err := dynSeq.GetAt(dynSeq.Len() - 1)
	if err != nil {
		t.Fatal(err)
	}

	if last != oldLen {
		t.Fatalf("the first value %v does not match what was inserted %v", last, oldLen)
	}

	getLast, err := dynSeq.GetLast()
	if err != nil {
		t.Fatal(err)
	}
	if getLast != oldLen {
		t.Fatalf("the first value %v does not match what was inserted %v", getLast, oldLen)
	}
}

func TestDeleteLast(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	startLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80, 10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}
	startCap := dynSeq.Cap()

	numOfDeleteFirstLoops := startLen / 2
	for range numOfDeleteFirstLoops {
		toBeDeleted, err := dynSeq.GetAt(dynSeq.Len() - 1)
		if err != nil {
			t.Fatal(err)
		}
		val, err := dynSeq.DeleteLast()
		if err != nil {
			t.Fatalf("unexpected error on DeleteLast: %e", err)
		}
		if toBeDeleted != val {
			t.Fatalf("The deleted value should match the value captured above: toBeDeleted %v <-> %v val", toBeDeleted, val)
		}
	}

	decLen := dynSeq.Len()
	if decLen != startLen-numOfDeleteFirstLoops {
		t.Fatalf("Length should have decreased by %v, curr len %v", numOfDeleteFirstLoops, dynSeq.Len())
	}
	decCap := dynSeq.Cap()
	if startCap/2 != decCap {
		t.Fatalf("starter capacity %v should have decreased by 2 folds, new capacity %v", startCap, decCap)
	}

	numOfDeleteSecondLoops := startLen / 4
	for range numOfDeleteSecondLoops {
		toBeDeleted, err := dynSeq.GetLast()
		if err != nil {
			t.Fatal(err)
		}
		val, err := dynSeq.DeleteLast()
		if err != nil {
			t.Fatalf("unexpected error on DeleteLast: %e", err)
		}
		if toBeDeleted != val {
			t.Fatalf("The deleted value should match the value captured above: toBeDeleted %v <-> %v val", toBeDeleted, val)
		}
	}

	secondDecLen := dynSeq.Len()
	if secondDecLen != decLen-numOfDeleteSecondLoops {
		t.Fatalf("Length should have decreased by %v, curr len %v", numOfDeleteSecondLoops, dynSeq.Len())
	}
	secondDecCap := dynSeq.Cap()
	if decCap/2 != secondDecCap {
		t.Fatalf("second capacity %v should have decreased by 2 folds, new capacity %v", startCap, decCap/2)
	}
}

func TestSetAt(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	_, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}

	idx := 2
	valToBeInserted := 11
	err = dynSeq.SetAt(idx, valToBeInserted)
	if err != nil {
		t.Fatalf("this index should be within the bonuds of the slice %v, %e", idx, err)
	}

	insertedVal, err := dynSeq.GetAt(idx)
	if err != nil {
		t.Fatal(err)
	}

	if valToBeInserted != insertedVal {
		t.Fatalf("Values do not match %v != %v", valToBeInserted, insertedVal)
	}

	outOfBoundsIdx := 11
	err = dynSeq.SetAt(outOfBoundsIdx, valToBeInserted)
	if err == nil {
		t.Fatalf("This index %v shoule be out of bounds", outOfBoundsIdx)
	}
}

func TestInsertAt(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	starterLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}

	idx := 2
	valToBeInserted := 11
	newLen, err := dynSeq.InsertAt(idx, valToBeInserted)
	if err != nil {
		t.Fatalf("this index should be within the bonuds of the slice %v, %e", idx, err)
	}

	val, err := dynSeq.GetAt(idx)
	if err != nil {
		t.Fatal(err)
	}

	if val != valToBeInserted {
		t.Fatalf("val %v does not mathc val to be inserted %v", val, valToBeInserted)
	}

	if starterLen >= newLen {
		t.Fatalf("Failed to inser at, length should be larger, starter %v new %v", starterLen, newLen)
	}
}

func TestInsertAtInLoop(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	starterLen, err := dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})
	if err != nil {
		t.Fatal(err)
	}

	idx := 2
	for i := range starterLen {
		_, err := dynSeq.InsertAt(idx, i+1)
		if err != nil {
			t.Fatalf("this index should be within the bonuds of the slice %v, %e", idx, err)
		}
	}

	val, err := dynSeq.GetAt(idx)
	if err != nil {
		t.Fatal(err)
	}
	if val != starterLen {
		t.Fatalf("val %v does not match val to be inserted %v", val, starterLen)
	}
}
