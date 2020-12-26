package linkedlistqueue

import "github.com/ravydv/go-data-structures/linkedlist/singlelinkedlist"

// Queue represent queue data structure
type Queue struct {
	queue *singlelinkedlist.SingleLinkedList
}

// New create new instance of queue
func New() *Queue {
	return &Queue{
		queue: &singlelinkedlist.SingleLinkedList{},
	}
}

// PushBack insert element in queue
func (q *Queue) PushBack(value interface{}) {
	q.queue.Add(value)
}

// Front return queue front element
func (q *Queue) Front() (value interface{}, ok bool) {
	value, ok = q.queue.Get(0)
	return
}

// PopFront remove queue top element
func (q *Queue) PopFront() (value interface{}, ok bool) {
	value, ok = q.queue.Get(0)
	q.queue.Remove(0)
	return
}

// Size return size of queue
func (q *Queue) Size() int {
	return q.queue.Size()
}

// Empty check if queue is empty or not
func (q *Queue) Empty() bool {
	return q.queue.Empty()
}

// Clear remove all element from the queue
func (q *Queue) Clear() {
	q.queue.Clear()
}

// String return string representation of queue
func (q *Queue) String() string {
	return q.queue.String()
}
