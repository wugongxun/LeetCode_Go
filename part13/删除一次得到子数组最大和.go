package main

import "math"

func maximumSum(arr []int) int {
	res, dp0, dp1 := math.MinInt/2, math.MinInt/2, math.MinInt/2
	for _, x := range arr {
		dp1 = max(dp1+x, dp0)
		dp0 = max(dp0, 0) + x
		res = max(res, dp0, dp1)
	}
	return res
}
