package dynamicsequence_test

import (
	"testing"

	dynamicsequence "ita/dynamicSequence"
)

func TestBuildAndGetAt(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})

	starterCap := dynSeq.Cap()
	starterLen := dynSeq.Len()

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
	dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})

	valToBeInserted := 1

	oldCap := dynSeq.Cap()
	oldLen := dynSeq.Len()
	newLen := dynSeq.InsertFirst(valToBeInserted)
	newCap := dynSeq.Cap()

	if oldLen >= newLen {
		t.Fatalf("new length should be larger than the new one. New length: %v", newLen)
	}

	if oldCap*2 != newCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old length -> %v", newCap, oldLen)
	}

	first, err := dynSeq.GetAt(0)
	if err != nil {
		t.Fatal(err)
	}

	if first != valToBeInserted {
		t.Fatalf("the first value %v does not match what was inserted %v", first, valToBeInserted)
	}

	for i := range dynSeq.Len() {
		dynSeq.InsertFirst(i)
	}
	seoondNewCap := dynSeq.Cap()

	if newCap*2 != seoondNewCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old length -> %v", newCap, oldLen)
	}
}

func TestDeleteFirstAndDownSize(t *testing.T) {
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80, 10, 20, 30, 40, 50, 60, 70, 80})
	startLen := dynSeq.Len()
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
	dynSeq.Build([]int{10, 20, 30, 40, 50, 60, 70, 80})

	valToBeInserted := 1

	oldCap := dynSeq.Cap()
	oldLen := dynSeq.Len()

	newLen := dynSeq.InsertLast(valToBeInserted)
	newCap := dynSeq.Cap()

	if oldLen >= newLen {
		t.Fatalf("new length should be larger than the new one. New length: %v", newLen)
	}

	if oldCap*2 != newCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old length -> %v", newCap, oldLen)
	}

	last, err := dynSeq.GetAt(dynSeq.Len() - 1)
	if err != nil {
		t.Fatal(err)
	}

	if last != valToBeInserted {
		t.Fatalf("the first value %v does not match what was inserted %v", last, valToBeInserted)
	}

	for i := range dynSeq.Len() {
		dynSeq.InsertFirst(i)
	}
	secondNewCap := dynSeq.Cap()

	if newCap*2 != secondNewCap {
		t.Fatalf("new capacity should be 6 times the old length: new capacity -> %v, old length -> %v", newCap, oldLen)
	}
}
