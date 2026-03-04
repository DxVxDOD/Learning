package main

import (
	"fmt"

	dynamicsequence "ita/dynamicSequence"
)

func main() {
	unSorted := []int{1, 3, 654, 765, 34, 34534, 5345, 345, 34534, 5346, 776876, 7634, 53}
	fmt.Println(unSorted)
	fmt.Println(InsertionSort(unSorted))

	seq := &Sequence[int]{}
	seq.Build([]int{10, 20, 30, 40, 50, 60, 70})
	seq.InsertAt(3, 69)
	fmt.Println(seq.IterSeq())
	seq.SetAt(5, 11)
	fmt.Println(seq.IterSeq())
	fmt.Println(seq.DeleteAt(3))
	fmt.Println(seq.IterSeq())
	fmt.Println(seq.DeleteFirst())
	fmt.Println(seq.IterSeq())
	seq.InsertFirst(69)
	fmt.Println(seq.IterSeq())
	fmt.Println(seq.DeleteLast())
	fmt.Println(seq.IterSeq())
	seq.InsertLast(69)
	fmt.Println(seq.IterSeq())
	dynSeq := &dynamicsequence.DynamicSequence[int]{}
	dynSeq.InsertFirst(1)
	fmt.Println(dynSeq.IterDynSeq())
}
