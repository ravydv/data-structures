package doublelinkedlist

import (
	"testing"
)

func TestDllNew(t *testing.T) {
	dll1 := New()
	if actualValue := dll1.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	dll2 := New(1, 2, 3, "4")
	if actualValue := dll2.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
}

func TestDllAdd(t *testing.T) {
	dll := New()
	dll.Add(1, 2, 3)
	if actualValue := dll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}
