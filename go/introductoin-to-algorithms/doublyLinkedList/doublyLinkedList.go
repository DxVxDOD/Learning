// Package doublylinkedlist represents a data structure called doubly linked lists. This data structure is sequential but not contiguous in memory.
package doublylinkedlist

import "errors"

type node[T any] struct {
	item T
	next *node[T]
	prev *node[T]
}

type DoublyLinkedList[T any] struct {
	lenght int
	head   *node[T]
	tail   *node[T]
}

func (d *DoublyLinkedList[T]) initWith(val T) {
	newNode := &node[T]{item: val}
	d.head = newNode
	d.tail = newNode
	d.lenght = 1
}

func (d *DoublyLinkedList[T]) InsertLast(val T) {
	if d.lenght == 0 {
		d.initWith(val)
	} else {
		oldTail := d.tail
		d.tail = &node[T]{item: val}
		oldTail.next = d.tail
		d.tail.prev = oldTail
		d.lenght++
	}
}

func (d *DoublyLinkedList[T]) Build(arr []T) {
	for _, v := range arr {
		d.InsertLast(v)
	}
}

func (d *DoublyLinkedList[T]) GetFirst() (T, error) {
	if d.lenght <= 0 {
		return *new(T), errors.New("cannot call GetFirst on a empty doubly linked list")
	}
	return d.head.item, nil
}

func (d *DoublyLinkedList[T]) GetLast() (T, error) {
	if d.lenght <= 0 {
		return *new(T), errors.New("cannot call GetLast on a empty doubly linked list")
	}
	return d.tail.item, nil
}

func (d *DoublyLinkedList[T]) GetAll() []T {
	arr := make([]T, d.lenght)
	curr := d.head
	i := 0
	for curr != nil {
		arr[i] = curr.item
		i++
		curr = curr.next
	}

	return arr
}
