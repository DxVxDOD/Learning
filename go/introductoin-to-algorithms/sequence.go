package main

import "errors"

type Sequencer[T any] interface {
	Build(x []T) int
	Len() int
	IterSeq() []T
	GetAt(idx int) T
	SetAt(idx int, val T)
	InsertAt(idx int, val T)
	DeleteAt(idx int) T
	InsertFirst(val T)
	DeleteFirst() T
	InsertLast(val T)
	DeleteLast() T
}

type Sequence[T any] struct {
	Seq    []T
	Length int
}

func (s *Sequence[T]) Build(x []T) int {
	s.Seq = x
	l := 0
	for range x {
		l++
	}
	s.Length = l
	return l
}

func (s *Sequence[T]) Len() int {
	return s.Length
}

func (s *Sequence[T]) IterSeq() []T {
	return s.Seq
}

func (s *Sequence[T]) GetAt(idx int) (T, error) {
	if idx >= s.Length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return s.Seq[idx], nil
}

func (s *Sequence[T]) SetAt(idx int, val T) error {
	if idx >= s.Length {
		return errors.New("index out of bounds")
	}
	s.Seq[idx] = val
	return nil
}

func (s *Sequence[T]) InsertAt(idx int, val T) error {
	if idx >= s.Length {
		return errors.New("index out of bounds")
	}
	newSeq := make([]T, s.Length+1)
	var temp T
	i := 0
	for ; i < s.Length; i++ {
		v := s.Seq[i]
		if i == idx {
			temp = v
			newSeq[i] = val
			i++
			newSeq[i] = temp
			break
		} else {
			newSeq[i] = v
		}
	}
	for ; i < s.Length+1; i++ {
		newSeq[i] = s.Seq[i-1]
	}
	s.Seq = newSeq
	s.Length = i
	return nil
}

func (s *Sequence[T]) DeleteAt(idx int) (T, error) {
	var zero T
	if idx >= s.Length {
		return zero, errors.New("index out of bounds")
	}
	deltVal := s.Seq[idx]
	for i := idx; i < s.Length-1; i++ {
		s.Seq[i] = s.Seq[i+1]
	}
	s.Seq[s.Length-1] = zero
	s.Length--
	return deltVal, nil
}

func (s *Sequence[T]) InsertFirst(val T) error {
	return s.InsertAt(0, val)
}

func (s *Sequence[T]) DeleteFirst() (T, error) {
	var zero T
	delVal, err := s.DeleteAt(0)
	if err != nil {
		return zero, err
	}
	return delVal, nil
}

func (s *Sequence[T]) InsertLast(val T) error {
	return s.InsertAt(s.Length-1, val)
}

func (s *Sequence[T]) DeleteLast() (T, error) {
	var zero T
	delVal, err := s.DeleteAt(s.Length - 1)
	if err != nil {
		return zero, err
	}
	return delVal, nil
}
