package main

import (
	"fmt"
	"tasks/ozon"
)

func main() {
	v := ozon.FindSubarrayWithMaxSum([]int{10, -3, -12, 8, 42, 1, -7, 0, 3})
	fmt.Println(v)
	fmt.Println(ozon.SlicesEqual(v[:], []int{3, 5}))
}
