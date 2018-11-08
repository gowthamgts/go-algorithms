package utils

import (
	"math/rand"
	"time"
)

func GetRandomArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n)
	}
	return array
}

func VerifySortedSet(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}

func Swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func MoveToStart(array []int, position int) []int {
	movedArray := make([]int, len(array))
	index := 0
	movedArray[index] = array[position]

	for i := 0; i < len(array); i++ {
		if i != position {
			movedArray[index+1] = array[i]
			index++
		}
	}

	return movedArray
}

func PickAndDrop(array []int, pickPosition int, dropPosition int) []int {

	if pickPosition <= dropPosition {
		return array
	}

	pdArray := make([]int, 0)
	pdArray = append(pdArray, array[:dropPosition]...)
	pdArray = append(pdArray, array[pickPosition])
	pdArray = append(pdArray, array[dropPosition:pickPosition]...)
	pdArray = append(pdArray, array[pickPosition+1:]...)

	return pdArray
}
