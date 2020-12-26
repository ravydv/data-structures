package queue

// Queue interface implemented by all queue
type Queue interface {
	PushBack(value interface{})
	Front() (value interface{}, ok bool)
	PopFront() (value interface{}, ok bool)
	Size() int
	Empty() bool
	Clear()
}
