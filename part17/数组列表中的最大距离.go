package main

import "math"

func maxDistance(arrays [][]int) int {
	res, mn, mx := 0, math.MaxInt/2, math.MinInt/2
	for _, arr := range arrays {
		x, y := arr[0], arr[len(arr)-1]
		res = max(res, y-mn, mx-x)
		mn = min(mn, x)
		mx = max(mx, y)
	}
	return res
}
