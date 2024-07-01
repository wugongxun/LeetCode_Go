package main

import "math"

func main() {
	print(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt / 2
	}
	for i, c := range cost {
		t := time[i] + 1
		for j := n; j > 0; j-- {
			dp[j] = min(dp[j], dp[max(j-t, 0)]+c)
		}
	}
	return dp[n]
}

//func paintWalls(cost []int, time []int) int {
//	n := len(cost)
//	cache := make([][]int, n)
//	for i := range cache {
//		cache[i] = make([]int, 2*n+1)
//		for j := range cache[i] {
//			cache[i][j] = -1
//		}
//	}
//	var dfs func(i, j int) int
//	dfs = func(i, j int) int {
//		if i < j {
//			return 0
//		}
//		if i < 0 {
//			return math.MaxInt / 2
//		}
//		p := &cache[i][j+n]
//		if *p != -1 {
//			return *p
//		}
//		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
//		return *p
//	}
//	return dfs(n-1, 0)
//}
