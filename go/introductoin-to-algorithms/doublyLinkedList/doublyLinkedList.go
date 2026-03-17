// Package doublylinkedlist represents a data structure called doubly linked lists. This data structure is sequential but not contiguous in memory.
package doublylinkedlist

import (
	"errors"
	"fmt"
)

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

func (d *DoublyLinkedList[T]) walkTo(idx int) *node[T] {
	var curr *node[T]
	if d.lenght/2 > idx {
		curr = d.tail
		for range idx {
			curr = curr.prev
		}
	} else {
		curr = d.head
		for range idx {
			curr = curr.next
		}
	}

	return curr
}

func (d *DoublyLinkedList[T]) Len() int {
	return d.lenght
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

func (d *DoublyLinkedList[T]) GetAt(idx int) (T, error) {
	if d.lenght <= 0 {
		return *new(T), errors.New("cannot call GetAt on a empty doubly linked list")
	}
	if idx < 0 || idx >= d.lenght {
		return *new(T), errors.New("index out of bounds")
	}
	return d.walkTo(idx).item, nil
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

func (d *DoublyLinkedList[T]) InsertFirst(val T) {
	if d.lenght == 0 {
		d.initWith(val)
	} else {
		oldHead := d.head
		d.head = &node[T]{item: val}
		d.head.next = oldHead
		oldHead.prev = d.head
		d.lenght++
	}
}

func (d *DoublyLinkedList[T]) InsertAt(val T, idx int) error {
	if d.lenght == 0 && idx == 0 {
		d.initWith(val)
		return nil
	}
	if idx < 0 || idx > d.lenght {
		return errors.New("index out of bounds")
	}

	if d.lenght == 0 {
		return fmt.Errorf("you can only insert at 0 on a empty doubly linked list %v", idx)
	}

	if idx == 0 {
		d.InsertFirst(val)
		return nil
	}
	if idx == d.lenght {
		d.InsertLast(val)
		return nil
	}

	nodeAtIdx := d.walkTo(idx)
	prevNode := nodeAtIdx.prev

	newNode := &node[T]{item: val}
	newNode.prev = prevNode
	newNode.next = nodeAtIdx

	nodeAtIdx.prev = newNode
	prevNode.next = newNode

	d.lenght++

	return nil
}

func (d *DoublyLinkedList[T]) Build(arr []T) {
	for _, v := range arr {
		d.InsertLast(v)
	}
}

func (d *DoublyLinkedList[T]) DeleteFirst() (T, error) {
	if d.lenght == 0 {
		return *new(T), errors.New("cannot delete from an empty list")
	}

	if d.lenght == 1 {
		item := d.head.item
		d.head = nil
		d.tail = nil
		return item, nil
	}

	nodeToBeDeleted := d.head
	d.head = nodeToBeDeleted.next
	d.head.prev = nil

	return nodeToBeDeleted.item, nil
}
