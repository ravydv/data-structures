package doublelinkedlist

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
