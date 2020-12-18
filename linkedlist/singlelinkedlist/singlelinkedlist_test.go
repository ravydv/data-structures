package singlelinkedlist

import (
	"testing"
)

func TestSllNew(t *testing.T) {
	sll1 := New()

	if actualValue := sll1.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	sll2 := New(1, "b")

	if actualValue := sll2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := sll2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := sll2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := sll2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestSllAdd(t *testing.T) {
	sll := New()
	sll.Add(1)
	sll.Add(2, 3)

	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := sll.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSllPrepend(t *testing.T) {
	sll := New()
	sll.Add(1)
	sll.Add(2, 3)

	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := sll.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	sll.Prepend(0)
	if actualValue := sll.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, ok := sll.Get(0); actualValue != 0 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSllRemove(t *testing.T) {
	sll := New()
	sll.Add(1)
	sll.Add(2, 3)

	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := sll.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	sll.Remove(0)
	if actualValue := sll.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := sll.Get(0); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestSllGet(t *testing.T) {
	sll := New()
	sll.Add(1)
	sll.Add(2, 3)

	if actualValue, ok := sll.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	sll.Prepend(0)

	if actualValue, ok := sll.Get(0); actualValue != 0 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	sll.Remove(2)
	if actualValue, ok := sll.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSllInsert(t *testing.T) {
	sll := New()
	sll.Add("c")
	sll.Insert(0, "a")
	sll.Insert(1, "b")
	sll.Insert(10, "x") // ignore
	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	sll.Insert(3, "d") // append
	if actualValue := sll.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := sll.String(), "a->b->c->d"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSllContains(t *testing.T) {
	sll := New()
	sll.Add("a")
	sll.Add("b", "c")
	if actualValue := sll.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := sll.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := sll.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	sll.Clear()
	if actualValue := sll.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := sll.Contains("a", "b", "c"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSllSet(t *testing.T) {
	sll := New()
	sll.Set(0, "a")
	sll.Set(1, "b")
	if actualValue := sll.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	sll.Set(2, "c") // append
	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	sll.Set(4, "d")  // ignore
	sll.Set(1, "bb") // update
	if actualValue := sll.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := sll.String(), "a->bb->c"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}
