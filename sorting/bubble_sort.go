package sorting

import "github.com/gowthamgts/go-algorithms"

func BubbleSort(array []int) []int {

	for j := 0; j < len(array); j++ {
		for i := 1; i < len(array) - j; i++ {
			if array[i-1] > array[i] {
				// swap
				utils.Swap(&array[i-1], &array[i])
			}
		}
	}

	return array
}
