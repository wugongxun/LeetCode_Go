package main

import "sort"

func minZeroArray(nums []int, queries [][]int) int {
	n := len(queries)
	diff := make([]int, len(nums)+1)
	res := sort.Search(n+1, func(k int) bool {
		clear(diff)
		for _, q := range queries[:k] {
			l, r, v := q[0], q[1], q[2]
			diff[l] += v
			diff[r+1] -= v
		}

		sumD := 0
		for i, x := range nums {
			sumD += diff[i]
			if x > sumD {
				return false
			}
		}
		return true
	})
	if res > n {
		return -1
	}
	return res
}

func main() {
	//print(minZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}))
	print(minZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}))
}
