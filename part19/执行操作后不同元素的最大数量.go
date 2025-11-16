package main

import (
	"math"
	"slices"
)

func maxDistinctElements(nums []int, k int) (res int) {
	slices.Sort(nums)
	pre := math.MinInt
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			res++
			pre = x
		}
	}
	return
}

func main() {
	println(maxDistinctElements([]int{7, 9, 10, 7, 7, 9, 8, 5, 10, 8}, 2))
}
