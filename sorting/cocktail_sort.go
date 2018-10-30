package sorting

import (
	"github.com/gowthamgts/go-algorithms"
)

func CocktailSort(array []int) []int {

	for i := 0; i < len(array)/2; i++ {

		left := 1
		right := len(array)

		for left <= right {

			if array[left-1] > array[left] {
				// swap
				utils.Swap(&array[left-1], &array[left])
			}
			left++

			if array[right-2] > array[right-1] {
				// swap
				utils.Swap(&array[right-2], &array[right-1])
			}
			right--
		}

	}

	return array
}
