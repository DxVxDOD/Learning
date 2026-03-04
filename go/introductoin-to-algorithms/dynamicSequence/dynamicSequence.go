// Package dynamicsequence respresents a srequnce of values which grows automatically while preserving order
package dynamicsequence

import (
	"errors"
)

type DynamicSequence[T any] struct {
	dynSeq   []T
	length   int
	capacity int
	start    int
	end      int
}

func (d *DynamicSequence[T]) Build(x []T) int {
	initLen := len(x)
	if initLen%2 != 0 {
		initLen++
	}
	starterCap := initLen * 3
	starterDynSeq := make([]T, starterCap)

	for i, v := range x {
		starterDynSeq[i+initLen] = v
	}

	d.length = initLen
	d.capacity = starterCap
	d.start = initLen
	d.end = initLen * 2

	d.dynSeq = starterDynSeq
	return d.length
}

func (d *DynamicSequence[T]) upSize() {
	oldLen := d.Len()
	if oldLen%2 != 0 {
		oldLen++
	}
	newLen := oldLen * 2
	newCap := newLen * 3
	newDynSeq := make([]T, newCap)

	for i, v := range d.dynSeq {
		newDynSeq[i+newLen] = v
	}

	d.capacity = newCap
	d.start = d.start + newLen
	d.end = d.end + newLen

	d.dynSeq = newDynSeq
}

func (d *DynamicSequence[T]) downSize() {
	oldCap := d.capacity
	newCap := oldCap / 2
	newDynSeq := make([]T, newCap)

	diffByHalf := newCap / 3
	for i := d.start; i < d.end; i++ {
		newDynSeq[diffByHalf] = d.dynSeq[i]
	}

	d.capacity = newCap
	d.start = diffByHalf
	d.end = diffByHalf * 2

	d.dynSeq = newDynSeq
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

	if d.length >= d.capacity/3 {
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

	if d.length < d.capacity/3 {
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

	if d.length >= d.capacity {
		d.upSize()
	}

	d.end++
	d.dynSeq[d.end] = val
	d.length++

	return d.length
}
