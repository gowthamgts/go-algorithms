package sorting

import (
	"github.com/gowthamgts/go-algorithms"
)

func InsertionSort(array []int) []int {

	for i := 1; i < len(array); i++ {
		dropPosition := i
		for j := i - 1; j >= 0; j-- {
			if array[i] < array[j] {
				// swap
				dropPosition = j
			}
		}
		array = utils.PickAndDrop(array, i, dropPosition)
	}

	return array
}
