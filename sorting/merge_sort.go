package sorting

func Merge(left []int, right []int) []int {

	mergedArray := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(mergedArray, right...)
		}

		if len(right) == 0 {
			return append(mergedArray, left...)
		}

		if left[0] <= right[0] {
			mergedArray = append(mergedArray, left[0])
			left = left[1:]
		} else {
			mergedArray = append(mergedArray, right[0])
			right = right[1:]
		}
	}

	return mergedArray
}

func MergeSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2

	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	return Merge(left, right)
}
