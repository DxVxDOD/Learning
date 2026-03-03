// Package dynamicsequence respresents a srequnce of values which grows automatically while preserving order
package dynamicsequence

import (
	"errors"
	"fmt"
)

type DynamicSequence[T any] struct {
	dynSeq   []T
	length   int
	capacity int
	start    int
	end      int
}

func (d *DynamicSequence[T]) Build(x []T) int {
	d.dynSeq = x
	d.length = len(x)
	d.capacity = cap(d.dynSeq)
	d.start = 0
	d.end = d.length
	return d.length
}

func (d *DynamicSequence[T]) upSize() {
	newDynSeq := make([]T, d.capacity*2+1)
	diffByHalf := (cap(newDynSeq) - cap(d.dynSeq) + 1) / 2

	for i, v := range d.dynSeq {
		newDynSeq[i+diffByHalf] = v
	}

	fmt.Println(newDynSeq)
	d.capacity = cap(newDynSeq)
	d.dynSeq = newDynSeq
	d.start = d.start + diffByHalf
	d.end = d.end + diffByHalf
}

func (d *DynamicSequence[T]) downSize() {
	newDynSeq := make([]T, d.capacity/2)
	diffByHalf := (cap(newDynSeq) + 1) / 2
	fmt.Println("diffByHalf: ", diffByHalf)
	for i := d.start; i < d.end; i++ {
		newDynSeq[i-diffByHalf] = d.dynSeq[i]
	}

	fmt.Println("newDynSeq: ", newDynSeq)
	d.capacity = cap(newDynSeq)
	d.dynSeq = newDynSeq
	d.start = d.start - diffByHalf
	d.end = d.end - diffByHalf
	fmt.Printf("newEnd: %v and new start: %v \n", d.end, d.start)
}

func (d *DynamicSequence[T]) Len() int {
	return d.length
}

func (d *DynamicSequence[T]) Cap() int {
	return d.capacity
}

func (d *DynamicSequence[T]) IterDynSeq() []T {
	return d.dynSeq[d.start:d.end]
}

func (d *DynamicSequence[T]) GetAt(idx int) (T, error) {
	if idx >= d.length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return d.dynSeq[idx+d.start], nil
}

func (d *DynamicSequence[T]) SetAt(idx int, val T) error {
	if idx >= d.length {
		return errors.New("index out of bounds")
	}
	d.dynSeq[idx] = val
	return nil
}

func (d *DynamicSequence[T]) InsertFirst(val T) int {
	if 0 >= d.length {
		d.Build([]T{val})
		return d.length
	}

	if d.length >= d.capacity {
		d.upSize()
	}

	d.start--
	d.dynSeq[d.start] = val
	d.length++

	return d.length
}

func (d *DynamicSequence[T]) DeleteFirst() (T, error) {
	if 0 >= d.length {
		var zero T
		return zero, errors.New("called delete on a empty sequence")
	}

	valToBeDeleted := d.dynSeq[d.start]

	if d.length <= d.capacity/4 {
		d.downSize()
	}

	d.start++
	d.length--

	return valToBeDeleted, nil
}

func (d *DynamicSequence[T]) InsertLast(val T) int {
	if 0 >= d.length {
		d.Build([]T{val})
		return d.length
	}

	fmt.Printf("len: %v, cap: %v \n", d.length, d.capacity)
	if d.length >= d.capacity {
		d.upSize()
	}

	d.end++
	d.dynSeq[d.end] = val
	d.length++

	return d.length
}
