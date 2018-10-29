package sorting

import "github.com/gowthamgts/go-algorithms"

func SelectionSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		minPosition := i

		for j := i + 1; j < len(array); j++ {
			if array[minPosition] > array [j] {
				// set the new position
				minPosition = j
			}
		}

		utils.Swap(&array[i], &array[minPosition])
	}

	return array
}
