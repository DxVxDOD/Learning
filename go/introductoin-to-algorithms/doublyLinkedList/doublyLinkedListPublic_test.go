package doublylinkedlist_test

import (
	"slices"
	"testing"

	doublylinkedlist "ita/doublyLinkedList"
)

func TestGetMethods(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})

	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}
	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}
	at, err := dl.GetAt(4)
	if err != nil {
		t.Fatal(err)
	}
	atEndish, err := dl.GetAt(5)
	if err != nil {
		t.Fatal(err)
	}

	if first != 1 {
		t.Fatalf("first value %v should be equal to 1", first)
	}

	if last != 8 {
		t.Fatalf("last value %v should be equal to 8", last)
	}

	if at != 5 {
		t.Fatalf("at value %v should be equal to 6", at)
	}

	if atEndish != 6 {
		t.Fatalf("atEndish %v should be 6", atEndish)
	}

	_, err = dl.GetAt(11)
	if err == nil {
		t.Fatalf("it should be index out of bounds")
	}

	_, err = dl.GetAt(-2)
	if err == nil {
		t.Fatalf("it should be index out of bounds")
	}
}

func TestEmptyGetMethods(t *testing.T) {
	emptyDL := doublylinkedlist.DoublyLinkedList[int]{}

	_, err := emptyDL.GetFirst()
	if err == nil {
		t.Fatal("GetFirst should err on an empty list")
	}

	_, err = emptyDL.GetLast()
	if err == nil {
		t.Fatal("GetLast should err on an empty list")
	}

	_, err = emptyDL.GetAt(1)
	if err == nil {
		t.Fatal("GetAt should err on an empty list")
	}
}

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
	at, err := dl.GetAt(4)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(all, first) {
		t.Fatalf("all should contain the first element %v", first)
	}

	if !slices.Contains(all, last) {
		t.Fatalf("all should contain the last element %v", last)
	}

	if !slices.Contains(all, at) {
		t.Fatalf("all should contain the at element %v", at)
	}
}

func TestInsertFirst(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1})
	starterLen := dl.Len()
	firstVal := 10

	dl.InsertFirst(10)
	afterLen := dl.Len()

	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	if afterLen <= starterLen {
		t.Fatalf("starterLen %v should be smaller than after len %v", starterLen, afterLen)
	}

	if afterLen != starterLen+1 {
		t.Fatalf("afterLen %v should be starterLen %v + 1", afterLen, starterLen)
	}

	if first != firstVal {
		t.Fatalf("get first value %v should equal the value inserted %v", first, firstVal)
	}

	emptyDL := doublylinkedlist.DoublyLinkedList[int]{}
	secondVal := 20
	emptyDL.InsertFirst(secondVal)

	secondFirst, err := emptyDL.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	if secondFirst != secondVal {
		t.Fatalf("second val %v should equal second get first %v", secondVal, secondFirst)
	}
}

func TestInsertLast(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1})
	starterLen := dl.Len()
	lastVal := 10

	dl.InsertLast(10)
	afterLen := dl.Len()

	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	if afterLen <= starterLen {
		t.Fatalf("starterLen %v should be smaller than after len %v", starterLen, afterLen)
	}

	if afterLen != starterLen+1 {
		t.Fatalf("afterLen %v should be starterLen %v + 1", afterLen, starterLen)
	}

	if last != lastVal {
		t.Fatalf("get last value %v should equal the value inserted %v", last, lastVal)
	}

	emptyDL := doublylinkedlist.DoublyLinkedList[int]{}
	secondVal := 20
	emptyDL.InsertLast(secondVal)

	secondLast, err := emptyDL.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	if secondLast != secondVal {
		t.Fatalf("second val %v should equal second get last %v", secondVal, secondLast)
	}
}

func TestInsertAt(t *testing.T) {
	dl := &doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})
	starteLen := dl.Len()
	atVal := 10
	atIdx := 4

	err := dl.InsertAt(atVal, atIdx)
	if err != nil {
		t.Fatal(err)
	}
	afterLen := dl.Len()

	if afterLen <= starteLen {
		t.Fatalf("starteLen %v should be smaller than afterLen %v", starteLen, afterLen)
	}

	if afterLen != starteLen+1 {
		t.Fatalf("starteLen %v should be only one smaller that afterLen %v", starteLen, afterLen)
	}

	getAtVal, err := dl.GetAt(atIdx)
	if err != nil {
		t.Fatal(err)
	}

	if getAtVal != atVal {
		t.Fatalf("getAtVal %v should equal the value inserted %v", getAtVal, atVal)
	}

	firstAtVal := 11
	err = dl.InsertAt(firstAtVal, 0)
	if err != nil {
		t.Fatal(err)
	}

	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	if first != firstAtVal {
		t.Fatalf("get first value %v should equal the inserted value %v", first, firstAtVal)
	}

	lastAtVal := 12
	err = dl.InsertAt(lastAtVal, dl.Len())
	if err != nil {
		t.Fatal(err)
	}

	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	if last != lastAtVal {
		t.Fatalf("get last value %v should equal the inserted value %v", last, lastAtVal)
	}
}

func TestDeleteFirst(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})
	starterLen := dl.Len()

	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	deletedVal, err := dl.DeleteFirst()
	if err != nil {
		t.Fatal(err)
	}

	afterLen := dl.Len()

	if afterLen >= starterLen {
		t.Fatalf("starterLen %v should be larger than afterLen %v", starterLen, afterLen)
	}
	if afterLen+1 != starterLen {
		t.Fatalf("afterLen %v + 1 should equal starterLen %v", afterLen, starterLen)
	}

	if first != deletedVal {
		t.Fatalf("the returned val by delete first %v should be equal the first val %v", deletedVal, first)
	}

	second, err := dl.GetAt(1)
	if err != nil {
		t.Fatal(err)
	}
	secondFirst, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	if second != secondFirst+1 {
		t.Fatalf("the second value %v should be larger by one per initialisation than the first %v", second, secondFirst)
	}
}

func TestInserAtEmpty(t *testing.T) {
	emptyDL := doublylinkedlist.DoublyLinkedList[int]{}

	emptyVal := 10
	emptyAtIdx := 11
	err := emptyDL.InsertAt(emptyVal, emptyAtIdx)
	if err == nil {
		t.Fatalf("You should not be able to insert above 0 on ane empty list, idx %v", emptyAtIdx)
	}

	err = emptyDL.InsertAt(emptyVal, 0)
	if err != nil {
		t.Fatal(err)
	}

	firstEmpty, err := emptyDL.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	if firstEmpty != emptyVal {
		t.Fatalf("the first value %v in the list should equal the inserted value %v", firstEmpty, emptyVal)
	}
}

func TestDeleteFirstOneElement(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1})

	first, err := dl.GetFirst()
	if err != nil {
		t.Fatal(err)
	}

	deletedVal, err := dl.DeleteFirst()
	if err != nil {
		t.Fatal(err)
	}

	if deletedVal != first {
		t.Fatalf("the returned val by delete first %v should be equal the first val %v", deletedVal, first)
	}

	_, err = dl.GetFirst()
	if err == nil {
		t.Fatal("the list should be epmpty")
	}

	if dl.Len() != 0 {
		t.Fatal("the lenght of the list should be 0")
	}
}

func TestDeleteFirstEmpty(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}

	_, err := dl.DeleteFirst()
	if err == nil {
		t.Fatal("should no be able to delete from an empty list")
	}
}

func TestDeleteLast(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1, 2, 3, 4, 5, 6, 7, 8})
	starterLen := dl.Len()

	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	deletedVal, err := dl.DeleteLast()
	if err != nil {
		t.Fatal(err)
	}

	afterLen := dl.Len()

	if afterLen >= starterLen {
		t.Fatalf("starterLen %v should be larger than afterLen %v", starterLen, afterLen)
	}
	if afterLen+1 != starterLen {
		t.Fatalf("afterLen %v + 1 should equal starterLen %v", afterLen, starterLen)
	}

	if last != deletedVal {
		t.Fatalf("the returned val by delete last %v should be equal the first val %v", deletedVal, last)
	}

	second, err := dl.GetAt(dl.Len() - 2)
	if err != nil {
		t.Fatal(err)
	}
	secondLast, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	if second+1 != secondLast {
		t.Fatalf("the second value %v should be smaller by one per initialisation than the last %v", second, secondLast)
	}
}

func TestDeleteLastOneElement(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}
	dl.Build([]int{1})

	last, err := dl.GetLast()
	if err != nil {
		t.Fatal(err)
	}

	deletedVal, err := dl.DeleteLast()
	if err != nil {
		t.Fatal(err)
	}

	if deletedVal != last {
		t.Fatalf("the returned val by delete last %v should be equal the last val %v", deletedVal, last)
	}

	_, err = dl.GetLast()
	if err == nil {
		t.Fatal("the list should be epmpty")
	}

	if dl.Len() != 0 {
		t.Fatal("the lenght of the list should be 0")
	}
}

func TestDeleteLastEmpty(t *testing.T) {
	dl := doublylinkedlist.DoublyLinkedList[int]{}

	_, err := dl.DeleteLast()
	if err == nil {
		t.Fatal("should no be able to delete from an empty list")
	}
}
