package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", maxSizeSlices([]int{1, 2, 3, 4, 5, 6}))
}

func maxSizeSlices(slices []int) int {
	return max(calculate(slices[1:]), calculate(slices[:len(slices)-1]))
}

func calculate(slices []int) int {
	N, n := len(slices), (len(slices)+1)/3
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MinInt
		}
	}
	dp[0][0], dp[0][1], dp[1][0], dp[1][1] = 0, slices[0], 0, max(slices[0], slices[1])
	for i := 2; i < N; i++ {
		dp[i][0] = 0
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i])
		}
	}
	return dp[N-1][n]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
