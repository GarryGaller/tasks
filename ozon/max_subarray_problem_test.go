package ozon

import (
    "testing"
)

type TestPairs struct {
    values   []int
    subarray [2]int
    sum      float64
}

var testsPairs = []TestPairs{
    {[]int{10, -3, -12, 8, 42, 1, -7, 0, 3}, [2]int{3, 5}, 51},
    {[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, [2]int{0, 8}, 45},
    {[]int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9}, [2]int{1, 9}, 45},
    {[]int{10, -3, 12, 8, 42, 1, -7, 0, 3}, [2]int{0, 5}, 70},
    {[]int{-7, 7, 12, 8, 42, 1, -7, 0, 3}, [2]int{1, 5}, 70},
    {[]int{1,2,3,4,5,-5}, [2]int{0, 4}, 15},
    {[]int{-1,-2,-3}, [2]int{0, 0}, -1},      // if the array contains only negative numbers, the maximum subarray would be the largest element itself. 
    {[]int{-3,-2,-1}, [2]int{2, 2}, -1},     //  if the array contains only negative numbers, the maximum subarray would be the largest element itself. 
}

func TestSubarray(t *testing.T) {
    for _, pair := range testsPairs {
        v := FindSubarrayWithMaxSum(pair.values)
        if v != pair.subarray {
            t.Error(
                "For", pair.values,
                "expected", pair.subarray,
                "got", v,
            )
        }
    }
}
