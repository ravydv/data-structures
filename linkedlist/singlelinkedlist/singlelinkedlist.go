package singlelinkedlist

import (
	"fmt"
	"strings"
)

// SingleLinkedList holds the node, where each node points to nex node
type SingleLinkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

// New create new single linked list and add the values passed, if any
func New(values ...interface{}) *SingleLinkedList {
	sll := &SingleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
	if len(values) > 0 {
		sll.Add(values...)
	}
	return sll
}

// Add appends one or more values at the end of the single linked list
func (sll *SingleLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &node{
			value: value,
			next:  nil,
		}
		if sll.size == 0 {
			sll.head = newNode
			sll.tail = newNode
		} else {
			sll.tail.next = newNode
			sll.tail = newNode
		}
		sll.size++
	}
}

// Prepend add one or more values at the start of the linked list
func (sll *SingleLinkedList) Prepend(values ...interface{}) {
	for _, v := range values {
		newNode := &node{
			value: v,
			next:  sll.head,
		}
		sll.head = newNode

		if sll.size == 0 {
			sll.tail = newNode
		}

		sll.size++
	}
}

// Get returns element at index
// Second return value is true if value exist at given index
func (sll *SingleLinkedList) Get(index int) (interface{}, bool) {

	if !sll.validateIndex(index) {
		return nil, false
	}

	current := sll.head
	for n := 0; n < index; n++ {
		current = current.next
	}

	return current.value, true
}

// Remove element at index form single linked list
func (sll *SingleLinkedList) Remove(index int) {
	if !sll.validateIndex(index) {
		return
	}
	if sll.size == 1 {
		sll.Clear()
		return
	}

	var previous *node

	current := sll.head

	for n := 0; n < index; n++ {
		previous = current
		current = current.next
	}

	if current == sll.head {
		sll.head = current.next
	} else if current == sll.tail {
		sll.tail = previous
	} else {
		previous.next = current.next
	}
	current = nil

	sll.size--
}

// Contains check if one or more values present in the set.
func (sll *SingleLinkedList) Contains(values ...interface{}) bool {

	for _, v := range values {
		found := false
		for current := sll.head; current != nil; current = current.next {
			if v == current.value {
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

// Insert values at given index, shift given index node to right, if any
func (sll *SingleLinkedList) Insert(index int, values ...interface{}) {
	if !sll.validateIndex(index) {
		if index == sll.size {
			sll.Add(values)
			return
		}
		return
	}
	sll.size += len(values)

	var previous *node
	current := sll.head
	for n := 0; n < index; n++ {
		previous = current
		current = current.next
	}

	if current == sll.head {
		oldFirstNode := sll.head
		for i, v := range values {
			newNode := &node{
				value: v,
				next:  nil,
			}

			if i == 0 {
				sll.head = newNode
			} else {
				previous.next = newNode
			}
			previous = newNode
		}
		previous.next = oldFirstNode
	} else {
		for _, v := range values {
			newNode := &node{
				value: v,
				next:  nil,
			}
			previous.next = newNode
			previous = newNode
		}
		previous.next = current
	}
}

// Set value at given index
// If index is equal to tail index it add that value to single linked list
func (sll *SingleLinkedList) Set(index int, value interface{}) {
	if !sll.validateIndex(index) {
		if index == sll.size {
			sll.Add(value)
		}
		return
	}

	current := sll.head
	for n := 0; n < index; n++ {
		current = current.next
	}
	current.value = value
}

// Empty return true if size of single linked list is zero, otherwise false
func (sll *SingleLinkedList) Empty() bool {
	return sll.size == 0
}

// Size return number of element in single linked list
func (sll *SingleLinkedList) Size() int {
	return sll.size
}

// Clear remove all element from the single linked list
func (sll *SingleLinkedList) Clear() {
	sll.head = nil
	sll.tail = nil
	sll.size = 0
}

// String return string representation of single linked list
func (sll *SingleLinkedList) String() string {
	values := []string{}

	for node := sll.head; node != nil; node = node.next {
		values = append(values, fmt.Sprintf("%v", node.value))
	}
	result := strings.Join(values, "->")
	return result
}

func (sll *SingleLinkedList) validateIndex(index int) bool {
	return index >= 0 && index < sll.size
}
