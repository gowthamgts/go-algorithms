package sorting

func BubbleSort(array []int) []int {
	temp := 0

	for j := 0; j < len(array); j++ {
		for i := 1; i < len(array) - j; i++ {
			if array[i-1] > array[i] {
				// swap
				temp = array[i-1]
				array[i-1] = array[i]
				array[i] = temp
			}
		}
	}

	return array
}
