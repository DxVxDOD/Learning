// Package linkedlist respresentes a data structure called linked lists. This data structure is sequential but not contiguous in memory.
package linkedlist

import (
	"errors"
)

type Linkedlist[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

type node[T any] struct {
	Item T
	next *node[T]
}

func (l *Linkedlist[T]) walkTo(idx int) (*node[T], error) {
	if idx >= l.length || idx < 0 {
		return nil, errors.New("index out of bounds")
	}

	nodeAtIdx := l.head
	i := 0

	for nodeAtIdx.next != nil && i < idx {
		nodeAtIdx = nodeAtIdx.next
		i++
	}
	return nodeAtIdx, nil
}

func (l *Linkedlist[T]) init(val T) {
	l.head = &node[T]{Item: val}
	l.tail = &node[T]{Item: val}
	l.length = 1
}

func (l *Linkedlist[T]) Build(arr []T) {
	// l.init(arr[0])

	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		l.InsertLast(arr[i])
	}
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

func (l *Linkedlist[T]) GetLast() (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot call GetLast on an empty Linkedlist")
	}

	return l.tail, nil
}

func (l *Linkedlist[T]) InsertLast(val T) {
	if l.tail == nil {
		l.init(val)
	} else {
		oldTail := l.tail
		l.tail = &node[T]{Item: val}
		oldTail.next = l.tail
		l.length++
	}
}

func (l *Linkedlist[T]) DeleteLast() (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	penultimNode, err := l.walkTo(l.length - 2)
	if err != nil {
		return nil, err
	}

	oldTail := penultimNode.next
	penultimNode.next = nil
	l.tail = penultimNode

	l.length--

	return oldTail, nil
}

func (l *Linkedlist[T]) GetFirst() (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot GetFirst on an empty linked list")
	}

	return l.head, nil
}

func (l *Linkedlist[T]) InsertFirst(val T) {
	if l.head == nil {
		l.init(val)
	} else {
		oldHead := l.head
		l.head = &node[T]{Item: val}
		l.head.next = oldHead

		l.length++
	}
}

func (l *Linkedlist[T]) DeleteFirst() (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	oldHead := l.head
	l.head = oldHead.next
	oldHead.next = nil

	l.length--

	return oldHead, nil
}

func (l *Linkedlist[T]) GetAt(idx int) (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot GetAt on an empty linked list")
	}

	return l.walkTo(idx)
}

func (l *Linkedlist[T]) InsertAt(val T, idx int) error {
	if l == nil {
		return errors.New("cannot insert at in an empty linked list")
	}

	if idx == l.length-1 {
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
	newNode := &node[T]{Item: val}
	newNode.next = nodeAtIdx
	nodeBeforeIdx.next = newNode

	l.length++

	return nil
}

func (l *Linkedlist[T]) DeleteAt(idx int) (*node[T], error) {
	if l == nil {
		return nil, errors.New("cannot delete from an empty linked list")
	}

	if idx == l.length-1 {
		return l.DeleteLast()
	}
	if idx == 0 {
		return l.DeleteFirst()
	}

	previousNode, err := l.walkTo(idx - 1)
	if err != nil {
		return nil, err
	}

	nodeToBeDeleted := previousNode.next
	nextNode := nodeToBeDeleted.next
	previousNode.next = nextNode

	l.length--

	return nodeToBeDeleted, nil
}
