package dqueue

import (
	"testing"
)

func TestPushFront(t *testing.T) {
	dq := New()
	if actualValue := dq.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushFront(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := dq.Front(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestPushBack(t *testing.T) {
	dq := New()
	if actualValue := dq.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushBack(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := dq.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestPopFront(t *testing.T) {
	dq := New()

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushFront(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	dq.PopFront()

	if actualValue := dq.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := dq.Front(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	dq.PopFront()

	if actualValue := dq.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := dq.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	dq.PopFront()

	if actualValue := dq.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	dq.PopFront() // no effect
}

func TestPopBack(t *testing.T) {
	dq := New()

	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushBack(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	dq.PopBack()

	if actualValue := dq.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := dq.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := dq.Back(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	dq.PopBack()

	if actualValue := dq.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := dq.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	dq.PopBack()

	if actualValue := dq.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	dq.PopBack() // no effect
}

func TestFront(t *testing.T) {
	dq := New()

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushFront(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := dq.Front(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	dq.PopFront()

	if actualValue := dq.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := dq.Front(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestBack(t *testing.T) {
	dq := New()

	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushBack(3)

	if actualValue := dq.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if actualValue, ok := dq.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := dq.Back(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	dq.PopFront()

	if actualValue := dq.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := dq.Front(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	dq.PopBack()

	if actualValue := dq.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := dq.Front(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
