package main

import "sort"

func main() {

}

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	res := make([]int64, len(queries))
	for i, query := range queries {
		index := sort.SearchInts(nums, query)
		s1 := query*index - prefix[index]
		s2 := prefix[n] - prefix[index] - query*(n-index)
		res[i] = int64(s1 + s2)
	}
	return res
}
