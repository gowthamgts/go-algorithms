package list

import (
	"testing"
)

func TestList_Append(t *testing.T) {

	list := New()

	for i := 0; i < 5; i++ {
		list.Append(i)
	}

	if list.isEmpty() {
		t.Error("List should not be empty")
	}

	for i := uint(0); i < 5; i++ {
		if list.Get(i).item != int(i) {
			t.Error("Item mismatch")
		}
	}
}

func TestList_Prepend(t *testing.T) {
	list := New()

	for i := 0; i < 5; i++ {
		list.Prepend(i)
	}

	if list.isEmpty() {
		t.Error("List should not be empty")
	}

	for i := uint(0); i < 5; i++ {
		if list.Get(i).item != (4 - int(i)) {
			t.Error("Item mismatch")
		}
	}
}
