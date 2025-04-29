package main

import "sort"

func largestDivisibleSubset(nums []int) (res []int) {
	sort.Ints(nums)
	n := len(nums)
	from := make([]int, n)
	dp := make([]int, n)
	maxI := 0
	for i := 0; i < n; i++ {
		from[i] = -1
		dp[i] = 0
	}

	for i, x := range nums {
		for j := 0; j < i; j++ {
			if x%nums[j] == 0 && dp[j] > dp[i] {
				dp[i] = dp[j]
				from[i] = j
			}
		}
		dp[i]++
		if dp[i] > dp[maxI] {
			maxI = i
		}
	}

	i := maxI
	for i >= 0 {
		res = append(res, nums[i])
		i = from[i]
	}
	return
}
