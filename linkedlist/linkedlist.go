package linkedlist

// LinkedList interface that all linkedlist implements
type LinkedList interface {
	Add(values ...interface{})
	Remove(index int)
	Contains(values ...interface{}) bool
	Get(index int) (interface{}, bool)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})
	Empty() bool
	Size() int
	Clear()
}
