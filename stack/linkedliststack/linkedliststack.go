package linkedliststack

import "github.com/ravydv/go-data-structures/linkedlist/singlelinkedlist"

// Stack repesent stack data structure
type Stack struct {
	stack *singlelinkedlist.SingleLinkedList
}

// New create new instance of the stack
func New() *Stack {
	return &Stack{
		stack: &singlelinkedlist.SingleLinkedList{},
	}
}

// Push insert element into stack
func (s *Stack) Push(value interface{}) {
	s.stack.Prepend(value)
}

// Pop remove top element from the stack
func (s *Stack) Pop() (value interface{}, ok bool) {
	value, ok = s.stack.Get(0)
	s.stack.Remove(0)
	return
}

// Peek return stack top element
func (s *Stack) Peek() (value interface{}, ok bool) {
	value, ok = s.stack.Get(0)
	return
}

// Size return number of element in the stack
func (s *Stack) Size() int {
	return s.stack.Size()
}

// Empty check emptyness of the stack
func (s *Stack) Empty() bool {
	return s.stack.Empty()
}

// Clear remove all element from stack
func (s *Stack) Clear() {
	s.stack.Clear()
}

// String return string representation of stack
func (s *Stack) String() string {
	return s.stack.String()
}
