package queue

import "testing"

func TestQueue(t *testing.T) {
	var queue = New()

	for i := 1; i < 6; i++ {
		queue.Enqueue(i)
	}

	if queue.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	if firstItem := queue.Peek(); firstItem.item != 1 {
		t.Error("Item not pushed", firstItem.item)
	}

	for i := 1; i < 6; i++ {
		if poppedElement := queue.Dequeue(); poppedElement != i {
			t.Error("Invalid Queue Variable", poppedElement)
		}
	}
}
