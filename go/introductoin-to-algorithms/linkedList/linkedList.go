// Package linkedlist respresentes a data structure called linked lists. This data structure is sequential but not contiguous in memory.
package linkedlist

import (
	"errors"
)

type Linkedlist[T any] struct {
	Item   T
	next   *Linkedlist[T]
	Head   *Linkedlist[T]
	length int
}

func (l *Linkedlist[T]) Build(arr []T) *Linkedlist[T] {
	var nextNode *Linkedlist[T]

	l.Item = arr[0]
	l.length = 1

	nextNode = l
	nextNode.Head = l

	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		nextNode = nextNode.InsertLast(arr[i])
	}

	return nextNode.Head
}

func (l *Linkedlist[T]) GetAll() []T {
	all := make([]T, l.length)
	nextNode := l.Head

	i := 0
	for nextNode != nil {
		all[i] = nextNode.Item
		i++
		nextNode = nextNode.next
	}

	return all
}

func (l *Linkedlist[T]) Len() int {
	return l.length
}

func (l *Linkedlist[T]) init(val T) *Linkedlist[T] {
	newNode := &Linkedlist[T]{Item: val}
	l = newNode
	l.Head = newNode
	l.length = 1
	return newNode
}

func (l *Linkedlist[T]) InsertLast(val T) *Linkedlist[T] {
	if l == nil {
		return l.init(val)
	}

	newNode := &Linkedlist[T]{Item: val}
	newNode.length = l.length + 1
	newNode.Head = l.Head
	newNode.Head.length = newNode.length

	lastNode := l.Head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
	return newNode
}

func (l *Linkedlist[T]) GetLast() *Linkedlist[T] {
	if l == nil {
		var zero T
		return l.init(zero)
	}

	lastNode := l.Head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	return lastNode
}

func (l *Linkedlist[T]) DeleteLast() (*Linkedlist[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	penultimNode := l.Head
	for penultimNode.next.next != nil {
		penultimNode = penultimNode.next
	}

	nodeToBeDeleted := penultimNode.next

	penultimNode.next = nil

	return nodeToBeDeleted, nil
}

// func (l *Linkedlist[T]) GetAt(idx int) (T, error) {
// 	var link *Linkedlist[T]
//
// 	for range idx {
// 		link = l.next
// 		if link == nil {
// 			var zero T
// 			return zero, errors.New("index out of range")
// 		}
// 	}
//
// 	return link.Item, nil
// }
