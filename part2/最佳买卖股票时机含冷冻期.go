package main

import "fmt"

func main() {
	fmt.Printf("%v", maxProfit([]int{1, 2, 3, 0, 2, 2, 4, 6, 8, 2, 5, 3}))
	//fmt.Printf("%v", maxProfit([]int{1}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	dp := make([][]int, n)
	dp[0] = append(dp[0], -prices[0], 0)
	dp[1] = append(dp[1], -min(prices[0], prices[1]), max(0, prices[1]-prices[0]))
	for i := 2; i < n; i++ {
		dp[i] = append(dp[i], max(dp[i-1][0], dp[i-2][1]-prices[i]))
		dp[i] = append(dp[i], max(dp[i-1][1], dp[i-1][0]+prices[i]))
	}
	return dp[n-1][1]
}
