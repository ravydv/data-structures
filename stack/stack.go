package stack

// Stack interface that implemented by all stack
type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	Size() int
	Empty() bool
	Clear()
}
