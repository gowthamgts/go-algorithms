package stack

import "testing"

func TestStackWithNumbers(t *testing.T) {
	var stack *Stack = New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("Test Fails for integers", i)
		}
	}
}


func TestStackWithStrings(t *testing.T) {
	var stringArray = []string {"Hello", "Dolly"}
	var stack *Stack = New()

	for i := 0; i < len(stringArray); i++ {
		stack.Push(stringArray[i])
	}

	for i := len(stringArray) - 1; i > 0; i-- {
		item := stack.Pop()
		if item != stringArray[i] {
			t.Error("Test Fails for strings", item)
		}
	}
}