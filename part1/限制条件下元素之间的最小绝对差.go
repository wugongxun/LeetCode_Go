package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Printf("%v", minAbsoluteDifference([]int{4, 3, 2, 4}, 2))
}

func minAbsoluteDifference(nums []int, x int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := math.MaxInt32
	t := make([]int, 0, n+2)
	t = append(t, math.MinInt32, math.MaxInt32)
	for i := x; i < n; i++ {
		insert := sort.SearchInts(t, nums[i-x])
		t = append(t[:insert], append([]int{nums[i-x]}, t[insert:]...)...)
		find := sort.SearchInts(t, nums[i])
		res = min(res, t[find]-nums[i])
		res = min(res, nums[i]-t[find-1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
