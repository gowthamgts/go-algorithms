package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)

	if a != 2 || b != 1 {
		t.Error("Swapping is wrong")
	}
}

func TestGetRandomArray(t *testing.T) {
	randomArray1 := GetRandomArray(10)
	randomArray2 := GetRandomArray(10)

	if len(randomArray1) < 10 || len(randomArray2) < 10 {
		t.Error("Length of random array is lower than expected")
	}

	var i int

	for i = 0; i < 10; i++ {
		if randomArray1[i] != randomArray2[i] {
			break
		}
	}

	if i == 10 {
		t.Error("Random array returns the same elements")
	}
}

func TestMoveToStart(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	array = MoveToStart(array, 2)
	assert.Equal(t, array, []int {3, 1, 2, 4, 5, 6, 7, 8, 9, 10})

	array = MoveToStart(array, 9)
	assert.Equal(t, array, []int {10, 3, 1, 2, 4, 5, 6, 7, 8, 9})

}

func TestPickAndDrop(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	array = PickAndDrop(array, 3, 0)
	assert.Equal(t, array, []int {4, 1, 2, 3, 5, 6, 7, 8, 9, 10})

	array = PickAndDrop(array, 9, 0)
	assert.Equal(t, array, []int {10, 4, 1, 2, 3, 5, 6, 7, 8, 9})

	array = PickAndDrop(array, 9, 1)
	assert.Equal(t, array, []int {10, 9, 4, 1, 2, 3, 5, 6, 7, 8})

	array = PickAndDrop(array, 6, 4)
	assert.Equal(t, array, []int {10, 9, 4, 1, 5, 2, 3, 6, 7, 8})

	// pickPosition smaller than dropPosition - should return the same array
	array = PickAndDrop(array, 6, 9)
	assert.Equal(t, array, []int {10, 9, 4, 1, 5, 2, 3, 6, 7, 8})

	array = PickAndDrop(array, 6, 6)
	assert.Equal(t, array, []int {10, 9, 4, 1, 5, 2, 3, 6, 7, 8})
}