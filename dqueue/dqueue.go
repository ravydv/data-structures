package dqueue

import (
	"github.com/ravydv/go-data-structures/linkedlist/doublelinkedlist"
)

// Dqueue data structure to reperesent dqueue
type Dqueue struct {
	dqueue *doublelinkedlist.DoubleLinkedList
}

// New create new instance of dqueue
func New() *Dqueue {
	return &Dqueue{
		dqueue: &doublelinkedlist.DoubleLinkedList{},
	}
}

// PushFront insert element in front of the dqueue
func (dq *Dqueue) PushFront(value interface{}) {
	dq.dqueue.Prepend(value)
}

// PushBack insert element at the back of the dqueue
func (dq *Dqueue) PushBack(value interface{}) {
	dq.dqueue.Append(value)
}

// Front return front element from the dqueue
func (dq *Dqueue) Front() (value interface{}, ok bool) {
	value, ok = dq.dqueue.Get(0)
	return
}

// Back return back element from the queue
func (dq *Dqueue) Back() (value interface{}, ok bool) {
	value, ok = dq.dqueue.Get(dq.Size() - 1)
	return
}

// PopFront remove front element of the dqueue
func (dq *Dqueue) PopFront() (value interface{}, ok bool) {
	value, ok = dq.dqueue.Get(0)
	dq.dqueue.Remove(0)
	return
}

// PopBack remove back element of the dqueue
func (dq *Dqueue) PopBack() (value interface{}, ok bool) {
	value, ok = dq.dqueue.Get(dq.Size() - 1)
	dq.dqueue.Remove(dq.Size() - 1)
	return
}

// Size return number of element in the dqueue
func (dq *Dqueue) Size() int {
	return dq.dqueue.Size()
}

// Empty check if dqueue is empty or not
func (dq *Dqueue) Empty() bool {
	return dq.dqueue.Empty()
}

// Clear remove all element from the dqueue
func (dq *Dqueue) Clear() {
	dq.dqueue.Clear()
}

// String return string representation of the dqueue
func (dq *Dqueue) String() string {
	return dq.dqueue.String()
}
