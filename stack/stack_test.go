package stack

import "testing"

func TestStackWithNumbers(t *testing.T) {
	var stack *Stack = New()

	if stack.IsEmpty() == false {
		t.Error("Stack should be empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	if stack.IsEmpty() == true {
		t.Error("Stack should not be empty")
	}

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("Test Fails for integers", i)
		}
	}
}

func TestStackWithStrings(t *testing.T) {
	var stringArray = []string{"Hello", "Dolly"}
	var stack *Stack = New()

	if stack.IsEmpty() == false {
		t.Error("Stack should be empty")
	}

	for i := 0; i < len(stringArray); i++ {
		stack.Push(stringArray[i])
	}

	if stack.IsEmpty() == true {
		t.Error("Stack should not be empty")
	}

	for i := len(stringArray) - 1; i > 0; i-- {
		item := stack.Pop()
		if item != stringArray[i] {
			t.Error("Test Fails for strings", item)
		}
	}
}
