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
