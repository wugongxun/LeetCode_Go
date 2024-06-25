package main

import "fmt"

func main() {
	fmt.Printf("%v", maxCoins([]int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3}))
}

func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

//func maxCoins(nums []int) int {
//	n := len(nums)
//	val := make([]int, n+2)
//	for i := 1; i <= n; i++ {
//		val[i] = nums[i-1]
//	}
//	val[0], val[n+1] = 1, 1
//	cache := make([][]int, n+2)
//	for i := 0; i < n+2; i++ {
//		cache[i] = make([]int, n+2)
//		for j := 0; j < n+2; j++ {
//			cache[i][j] = -1
//		}
//	}
//	return solve(0, n+1, val, cache)
//}
//
//func solve(l, r int, val []int, cache [][]int) int {
//	if l >= r-1 {
//		return 0
//	}
//	if cache[l][r] != -1 {
//		return cache[l][r]
//	}
//	for i := l + 1; i < r; i++ {
//		sum := val[l] * val[i] * val[r]
//		sum += solve(l, i, val, cache) + solve(i, r, val, cache)
//		cache[l][r] = max(cache[l][r], sum)
//	}
//	return cache[l][r]
//}
