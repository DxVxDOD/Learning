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

func (d *DynamicSequence[T]) Build(x []T) (int, error) {
	initLen := len(x)

	if initLen == 0 {
		return 0, errors.New("cannot build from an empty sequence")
	}

	if initLen%2 != 0 {
		initLen++
	}

	starterCap := initLen * 3
	starterDynSeq := make([]T, starterCap)

	for i, v := range x {
		starterDynSeq[i+initLen] = v
	}

	d.length = len(x)
	d.capacity = starterCap
	d.start = initLen
	d.end = initLen * 2

	d.dynSeq = starterDynSeq

	return d.length, nil
}

func (d *DynamicSequence[T]) upSize() {
	oldCap := d.capacity
	newCap := oldCap * 2
	newDynSeq := make([]T, newCap)

	j := d.length
	for i := d.start; i < d.end; i++ {
		newDynSeq[j] = d.dynSeq[i]
		j++
	}

	d.capacity = newCap
	d.start = d.length
	d.end = j

	d.dynSeq = newDynSeq
}

func (d *DynamicSequence[T]) downSize() {
	oldCap := d.capacity
	newCap := oldCap / 2
	newDynSeq := make([]T, newCap)

	diffByHalf := newCap / 3
	j := diffByHalf
	for i := d.start; i < d.end; i++ {
		newDynSeq[j] = d.dynSeq[i]
		j++
	}

	d.capacity = newCap
	d.start = diffByHalf
	d.end = j

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

func (d *DynamicSequence[T]) SetAt(idx int, val T) error {
	if idx >= d.length {
		return errors.New("index out of bounds")
	}

	d.dynSeq[d.start+idx] = val

	return nil
}

func (d *DynamicSequence[T]) GetAt(idx int) (T, error) {
	if idx >= d.length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return d.dynSeq[idx+d.start], nil
}

func (d *DynamicSequence[T]) GetFirst() (T, error) {
	if d.length <= 0 {
		var zero T
		return zero, errors.New("cannot call GetFirst on an empty sequence")
	}
	return d.dynSeq[d.start], nil
}

func (d *DynamicSequence[T]) GetLast() (T, error) {
	if d.length <= 0 {
		var zero T
		return zero, errors.New("cannot call GetLast on an empty sequence")
	}
	return d.dynSeq[d.end-1], nil
}

func (d *DynamicSequence[T]) InsertFirst(val T) (int, error) {
	if 0 >= d.length {
		l, err := d.Build([]T{val})
		if err != nil {
			return 0, err
		}
		return l, nil
	}

	d.length++

	if d.start <= 1 {
		d.upSize()
	}

	d.start--
	d.dynSeq[d.start] = val

	return d.length, nil
}

func (d *DynamicSequence[T]) DeleteFirst() (T, error) {
	var zero T
	if 0 >= d.length {
		return zero, errors.New("called delete on a empty sequence")
	}

	valToBeDeleted := d.dynSeq[d.start]
	d.dynSeq[d.start] = zero

	d.length--
	d.start++

	if d.length <= d.capacity/6 {
		d.downSize()
	}

	return valToBeDeleted, nil
}

func (d *DynamicSequence[T]) InsertLast(val T) (int, error) {
	if 0 >= d.length {
		l, err := d.Build([]T{val})
		if err != nil {
			return 0, err
		}
		return l, nil
	}

	d.length++
	if d.end >= d.capacity-1 {
		d.upSize()
	}

	d.dynSeq[d.end] = val
	d.end++

	return d.length, nil
}

func (d *DynamicSequence[T]) DeleteLast() (T, error) {
	var zero T
	if 0 >= d.length {
		return zero, errors.New("called delete on a empty sequence")
	}

	valToBeDeleted := d.dynSeq[d.end-1]
	d.dynSeq[d.end-1] = zero

	d.length--
	d.end--

	if d.length <= d.capacity/6 {
		d.downSize()
	}

	return valToBeDeleted, nil
}

func (d *DynamicSequence[T]) InsertAt(idx int, val T) (int, error) {
	if 0 >= d.length {
		l, err := d.Build([]T{val})
		if err != nil {
			return 0, err
		}
		return l, nil
	}

	if idx >= d.length || idx < 0 {
		return 0, errors.New("index out of bounds")
	}

	d.length++

	if d.end >= d.capacity-1 {
		d.upSize()
	}

	realIdx := idx + d.start
	t1 := d.dynSeq[realIdx]
	d.dynSeq[realIdx] = val

	d.end++
	for i := realIdx + 1; i < d.end; i++ {
		t2 := d.dynSeq[i]
		d.dynSeq[i] = t1
		t1 = t2
	}

	return d.length, nil
}
