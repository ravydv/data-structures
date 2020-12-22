package doublelinkedlist

import (
	"fmt"
	"strings"
)

// DoubleLinkedList  holds the linked list node, where each node points to next node and previous node
type DoubleLinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
	prev  *node
}

// New create new double linked list and add the values passed, if any
func New(values ...interface{}) *DoubleLinkedList {
	dll := &DoubleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}

	if len(values) > 0 {
		for _, v := range values {
			dll.Add(v)
		}
	}
	return dll
}

// Add appends one or more values at the end of the linked list
func (dll *DoubleLinkedList) Add(values ...interface{}) {
	for _, v := range values {
		newNode := &node{
			value: v,
			next:  nil,
			prev:  nil,
		}

		if dll.size == 0 {
			dll.head = newNode
			dll.tail = newNode
		} else {
			dll.tail.next = newNode
			newNode.prev = dll.tail
			dll.tail = newNode
		}
		dll.size++
	}
}

// Append one or more values at end of the linked list
func (dll *DoubleLinkedList) Append(values ...interface{}) {
	for _, v := range values {
		dll.Add(v)
	}
}

// Prepend add one or more values at the start of the linked list
func (dll *DoubleLinkedList) Prepend(values ...interface{}) {
	for _, v := range values {
		newNode := &node{
			value: v,
			next:  nil,
			prev:  nil,
		}

		if dll.size == 0 {
			dll.head = newNode
			dll.tail = newNode
		} else {
			dll.head.prev = newNode
			dll.head = newNode
		}
		dll.size++
	}
}

// Insert value(s) at specified index, shift any node to right at specified index, if exist
func (dll *DoubleLinkedList) Insert(index int, values ...interface{}) {

	if !dll.validateIndex(index) {
		if index == dll.size {
			for _, v := range values {
				dll.Add(v)
			}
			return
		}
		return
	}

	dll.size += len(values)

	var previous *node
	current := dll.head

	for i := 0; i < index; i++ {
		previous = current
		current = current.next
	}

	if current == dll.head {
		oldHead := dll.head
		var newHead *node
		var newTail *node
		for i, v := range values {
			newNode := &node{
				value: v,
				next:  nil,
				prev:  nil,
			}
			if i == 0 {
				newHead = newNode
				newTail = newNode
			} else {
				newTail.next = newNode
				newNode.prev = newTail
				newTail = newNode
			}
		}
		newTail.next = oldHead
		oldHead.prev = newTail
		dll.head = newHead
	} else {
		for _, v := range values {
			newNode := &node{
				value: v,
				next:  nil,
				prev:  nil,
			}

			previous.next = newNode
			newNode.prev = previous
			previous = newNode
		}
		previous.next = current
		current.prev = previous
	}
}

// Get returns value at specified index.
// Returs nil, false if index is invalid
func (dll *DoubleLinkedList) Get(index int) (interface{}, bool) {
	if !dll.validateIndex(index) {
		return nil, false
	}

	var current *node = dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, true
}

// Contains returns true if all values are present in the linked list, otherwisae false
func (dll *DoubleLinkedList) Contains(values ...interface{}) bool {
	for _, v := range values {
		found := false
		for current := dll.head; current != nil; current = current.next {
			if current.value == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Set value at specified index.
// if index equals to size of linked list, value is added at end of the list.
func (dll *DoubleLinkedList) Set(index int, value interface{}) {
	if !dll.validateIndex(index) {
		if dll.size == index {
			dll.Add(value)
			return
		}
		return
	}

	currrent := dll.head
	for i := 0; i < index; i++ {
		currrent = currrent.next
	}
	currrent.value = value
}

// Remove element at specified index
func (dll *DoubleLinkedList) Remove(index int) {
	if !dll.validateIndex(index) {
		return
	}

	if dll.size == 1 {
		dll.Clear()
		return
	}

	var previous *node
	current := dll.head
	for i := 0; i < index; i++ {
		previous = current
		current = current.next
	}

	if current == dll.head {
		dll.head = current.next
	} else if current == dll.tail {
		dll.tail = previous
	} else {
		previous.next = current.next
		current.next.prev = previous
	}
	dll.size--
	current = nil
}

// Empty check emptyness of linked list
func (dll *DoubleLinkedList) Empty() bool {
	return dll.size == 0
}

// Size return number of element in the linked list
func (dll *DoubleLinkedList) Size() int {
	return dll.size
}

// Clear remove all the element from the linked list
func (dll *DoubleLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

// String return string representation of linked list
func (dll *DoubleLinkedList) String() string {
	result := []string{}

	for current := dll.head; current != nil; current = current.next {
		result = append(result, fmt.Sprintf("%v", current.value))
	}
	return strings.Join(result, ", ")
}

func (dll *DoubleLinkedList) validateIndex(index int) bool {
	return index >= 0 && index < dll.size
}
