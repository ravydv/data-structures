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

	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestDllGet(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestDllContains(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	list.Clear()
	if actualValue := list.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Contains("a", "b", "c"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestDllInsert(t *testing.T) {
	list := New()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := list.String(), "a, b, c, d"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestDllSet(t *testing.T) {
	list := New()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := list.String(), "a, bb, c"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	list.Set(2, "cc")
	list.Set(0, "aa")
	if actualValue, expectedValue := list.String(), "aa, bb, cc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestDllRemove(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}
