package main

import (
	"errors"
)

type DynamicSequence[T any] struct {
	DynSeq   []T
	Length   int
	Capacity int
}

func (d *DynamicSequence[T]) Build(x []T) int {
	l := 0
	for range x {
		l++
	}
	d.Length = l
	d.Capacity = l * 2
	d.DynSeq = make([]T, d.Length, d.Capacity)
	return l
}

func (d *DynamicSequence[T]) IterDynSeq() []T {
	return d.DynSeq
}

func (d *DynamicSequence[T]) GetAt(idx int) (T, error) {
	if idx >= d.Length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return d.DynSeq[idx], nil
}

func (d *DynamicSequence[T]) SetAt(idx int, val T) error {
	if idx >= d.Length {
		return errors.New("index out of bounds")
	}
	d.DynSeq[idx] = val
	return nil
}
