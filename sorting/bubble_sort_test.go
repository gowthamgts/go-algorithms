package sorting

import (
	"fmt"
	"github.com/gowthamgts/go-algorithms"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	randomArray := utils.GetRandomArray(10)
	fmt.Println("Sorting Random Array: ", randomArray)

	sortedArray := BubbleSort(randomArray)
	fmt.Println("Sorted Array: ", sortedArray)

	properlySorted := utils.VerifySortedSet(sortedArray)

	if properlySorted {
		fmt.Println("Sorting Verified!")
	} else {
		t.Error("Sorting Verification Failed")
	}
}
