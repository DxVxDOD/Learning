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

func (d *DynamicSequence[T]) resize() {
	newDynSeq := make([]T, d.capacity*2)
	diffByHalf := (cap(newDynSeq) - cap(d.dynSeq) + 1) / 2

	for i, v := range d.dynSeq {
		newDynSeq[i+diffByHalf] = v
	}

	fmt.Println(diffByHalf, d.start, d.end)

	d.capacity = cap(newDynSeq)
	d.dynSeq = newDynSeq
	d.start = d.start + diffByHalf
	d.end = d.end + diffByHalf

	fmt.Println("Second time around", diffByHalf, d.start, d.end)
}

func (d *DynamicSequence[T]) Len() int {
	return d.length
}

func (d *DynamicSequence[T]) Cap() int {
	return d.capacity
}

func (d *DynamicSequence[T]) IterDynSeq() []T {
	fmt.Println(d.start, d.end)
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
		d.resize()
	}

	d.start--
	d.dynSeq[d.start] = val
	d.length++
	return d.length
}
