package main

import "math"

func maxProfit(k int, prices []int) int {
	dp := make([][2]int, k+2)
	for i := 1; i <= k+1; i++ {
		dp[i][1] = math.MinInt / 2
	}
	dp[0][0] = math.MinInt / 2
	for _, price := range prices {
		for j := 1; j <= k+1; j++ {
			dp[j][0] = max(dp[j][0], dp[j][1]+price)
			dp[j][1] = max(dp[j][1], dp[j-1][0]-price)
		}
	}
	return dp[k+1][0]

	//n := len(prices)
	//dp := make([][][2]int, n+1)
	//for i := range dp {
	//	dp[i] = make([][2]int, k+2)
	//	for j := range dp[i] {
	//		dp[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
	//	}
	//}
	//for i := 1; i <= k+1; i++ {
	//	dp[0][i][0] = 0
	//}
	//for i, price := range prices {
	//	for j := 1; j <= k+1; j++ {
	//		dp[i+1][j][0] = max(dp[i][j][0], dp[i][j][1]+price)
	//		dp[i+1][j][1] = max(dp[i][j][1], dp[i][j-1][0]-price)
	//	}
	//}
	//return dp[n][k+1][0]
}
