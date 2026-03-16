// Package linkedlist represents a data structure called linked lists. This data structure is sequential but not contiguous in memory.
package linkedlist

import (
	"errors"
)

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

type node[T any] struct {
	item T
	next *node[T]
}

func (l *LinkedList[T]) walkTo(idx int) (*node[T], error) {
	if idx >= l.length || idx < 0 {
		return nil, errors.New("index out of bounds")
	}

	nodeAtIdx := l.head
	i := 0

	for i < idx {
		nodeAtIdx = nodeAtIdx.next
		i++
	}
	return nodeAtIdx, nil
}

func (l *LinkedList[T]) initWith(val T) {
	l.head = &node[T]{item: val}
	l.tail = l.head
	l.length = 1
}

func (l *LinkedList[T]) Build(arr []T) {
	for _, v := range arr {
		l.InsertLast(v)
	}
}

func (l *LinkedList[T]) GetAll() []T {
	all := make([]T, l.length)
	nextNode := l.head

	i := 0
	for nextNode != nil {
		all[i] = nextNode.item
		i++
		nextNode = nextNode.next
	}

	return all
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) GetLast() (T, error) {
	if l.tail == nil {
		return *new(T), errors.New("cannot call GetLast on an empty Linkedlist")
	}

	return l.tail.item, nil
}

func (l *LinkedList[T]) InsertLast(val T) {
	if l.tail == nil {
		l.initWith(val)
	} else {
		oldTail := l.tail
		l.tail = &node[T]{item: val}
		oldTail.next = l.tail
		l.length++
	}
}

func (l *LinkedList[T]) DeleteLast() (T, error) {
	if l.tail == nil {
		return *new(T), errors.New("cannot delete from an empty linked list")
	}

	if l.length == 1 {
		nodeToBeDeleted := l.tail
		l.tail = nil
		l.head = nil
		l.length--
		return nodeToBeDeleted.item, nil
	}

	penultimNode, err := l.walkTo(l.length - 2)
	if err != nil {
		return *new(T), err
	}

	oldTail := penultimNode.next
	penultimNode.next = nil
	l.tail = penultimNode

	l.length--

	return oldTail.item, nil
}

func (l *LinkedList[T]) GetFirst() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot GetFirst on an empty linked list")
	}

	return l.head.item, nil
}

func (l *LinkedList[T]) InsertFirst(val T) {
	if l.head == nil {
		l.initWith(val)
	} else {
		oldHead := l.head
		l.head = &node[T]{item: val}
		l.head.next = oldHead

		l.length++
	}
}

func (l *LinkedList[T]) DeleteFirst() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot delete from an empty linked list")
	}

	if l.length == 1 {
		nodeToBeDeleted := l.head
		l.tail = nil
		l.head = nil
		l.length--
		return nodeToBeDeleted.item, nil
	}

	oldHead := l.head
	l.head = oldHead.next
	oldHead.next = nil

	l.length--

	return oldHead.item, nil
}

func (l *LinkedList[T]) GetAt(idx int) (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot GetAt on an empty linked list")
	}

	if idx < 0 || idx >= l.length {
		return *new(T), errors.New("index out of bounds")
	}

	nodeAtIdx, err := l.walkTo(idx)
	if err != nil {
		return *new(T), err
	}

	return nodeAtIdx.item, nil
}

func (l *LinkedList[T]) InsertAt(val T, idx int) error {
	if idx < 0 || idx > l.length {
		return errors.New("index out of bounds")
	}

	if l.head == nil {
		l.initWith(val)
		return nil
	}

	if idx == l.length {
		l.InsertLast(val)
		return nil
	}

	if idx == 0 {
		l.InsertFirst(val)
		return nil
	}

	nodeBeforeIdx, err := l.walkTo(idx - 1)
	if err != nil {
		return err
	}

	nodeAtIdx := nodeBeforeIdx.next
	newNode := &node[T]{item: val}
	newNode.next = nodeAtIdx
	nodeBeforeIdx.next = newNode

	l.length++

	return nil
}

func (l *LinkedList[T]) DeleteAt(idx int) (T, error) {
	if l.head == nil {
		return *new(T), errors.New("cannot delete from an empty linked list")
	}

	if idx < 0 || idx >= l.length {
		return *new(T), errors.New("index out of bounds")
	}

	if idx == l.length-1 {
		return l.DeleteLast()
	}
	if idx == 0 {
		return l.DeleteFirst()
	}

	previousNode, err := l.walkTo(idx - 1)
	if err != nil {
		return *new(T), err
	}

	nodeToBeDeleted := previousNode.next
	nextNode := nodeToBeDeleted.next
	previousNode.next = nextNode
	nodeToBeDeleted.next = nil

	l.length--

	return nodeToBeDeleted.item, nil
}
