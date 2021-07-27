package ozon

import (
	"fmt"
)

func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FindSubarrayWithMaxSum(arr []int) [2]int {
	startSum := arr[0]
	leftBound, rightBound, partialSum := 0, 0, 0
	minusPos := -1
	n := len(arr)

	for r := 0; r < n; r++ {
		partialSum += arr[r]

		if partialSum > startSum {
			startSum = partialSum
			leftBound = minusPos + 1
			rightBound = r
		}

		if partialSum < 0 {
			partialSum = 0
			minusPos = r
		}
	}
	fmt.Println(startSum)
	return [2]int{leftBound, rightBound}
}
