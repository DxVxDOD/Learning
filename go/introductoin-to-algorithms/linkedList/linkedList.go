// Package linkedlist respresentes a data structure called linked lists. This data structure is sequential but not contiguous in memory.
package linkedlist

import (
	"errors"
)

type Linkedlist[T any] struct {
	Item   T
	next   *Linkedlist[T]
	head   *Linkedlist[T]
	length int
}

func (l *Linkedlist[T]) walkFromStartToEnd() *Linkedlist[T] {
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	return lastNode
}

func (l *Linkedlist[T]) init(val T) *Linkedlist[T] {
	newNode := &Linkedlist[T]{Item: val}

	l = newNode
	l.head = newNode
	l.length = 1

	return newNode
}

func (l *Linkedlist[T]) Build(arr []T) *Linkedlist[T] {
	nextNode := l.init(arr[0])
	nextNode.head = nextNode

	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		nextNode = nextNode.InsertLast(arr[i])
	}

	return nextNode.head
}

func (l *Linkedlist[T]) GetAll() []T {
	all := make([]T, l.length)
	nextNode := l.head

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

func (l *Linkedlist[T]) GetLast() *Linkedlist[T] {
	if l == nil {
		var zero T
		return l.init(zero)
	}

	lastNode := l.walkFromStartToEnd()

	return lastNode
}

func (l *Linkedlist[T]) InsertLast(val T) *Linkedlist[T] {
	if l == nil {
		return l.init(val)
	}

	l.length++
	l.head.length = l.length

	newNode := &Linkedlist[T]{Item: val}
	newNode.length = l.length
	newNode.head = l.head

	lastNode := l.walkFromStartToEnd()
	lastNode.next = newNode

	return newNode
}

func (l *Linkedlist[T]) DeleteLast() (*Linkedlist[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	penultimNode := l.head
	for penultimNode.next.next != nil {
		penultimNode = penultimNode.next
	}

	nodeToBeDeleted := penultimNode.next

	penultimNode.next = nil

	l.length--
	l.head.length = l.length

	return nodeToBeDeleted, nil
}

func (l *Linkedlist[T]) GetFirst() (*Linkedlist[T], error) {
	if l == nil {
		return nil, errors.New("cannot GetFirst on an empty linked list")
	}

	return l.head, nil
}

func (l *Linkedlist[T]) InsertFirst(val T) *Linkedlist[T] {
	if l == nil {
		return l.init(val)
	}

	newHead := &Linkedlist[T]{Item: val}
	secondNode := l.head
	newHead.next = secondNode
	newHead.head = newHead

	l.head = newHead

	l.length++
	l.head.length = l.length

	return newHead
}

func (l *Linkedlist[T]) DeleteFirst() (*Linkedlist[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	nodeToBeDeleted := l.head
	secondNode := nodeToBeDeleted.next

	l.head = secondNode
	l.length--
	l.head.length = l.length

	return nodeToBeDeleted, nil
}
