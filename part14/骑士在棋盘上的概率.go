package main

import "fmt"

var dirs = []struct{ x, y int }{{1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, n+4)
	for i := range dp {
		dp[i] = make([][]float64, n+4)
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+1)
		}
	}
	for i := 2; i < n+2; i++ {
		for j := 2; j < n+2; j++ {
			dp[i][j][0] = 1
		}
	}
	for step := 1; step < k+1; step++ {
		for i := 2; i < n+2; i++ {
			for j := 2; j < n+2; j++ {
				for _, d := range dirs {
					dp[i][j][step] += dp[i+d.x][j+d.y][step-1] / 8
				}
			}
		}
	}
	return dp[row+2][column+2][k]
}

//func knightProbability(n int, k int, row int, column int) float64 {
//	cache := make([][][]float64, n)
//	for i := range cache {
//		cache[i] = make([][]float64, n)
//		for j := range cache[i] {
//			cache[i][j] = make([]float64, k+1)
//			for step := range cache[i][j] {
//				cache[i][j][step] = -1
//			}
//		}
//	}
//	var dfs func(i, j, step int) float64
//	dfs = func(i, j, step int) float64 {
//		if step == 0 {
//			return 1
//		}
//		p := &cache[i][j][step]
//		if *p != -1 {
//			return cache[i][j][step]
//		}
//		inRate := float64(0)
//		for _, d := range dirs {
//			if i+d.x >= 0 && i+d.x < n && j+d.y >= 0 && j+d.y < n {
//				inRate += dfs(i+d.x, j+d.y, step-1) / 8
//			}
//		}
//		*p = inRate
//		return inRate
//	}
//	return dfs(row, column, k)
//}

func main() {
	//fmt.Printf("%v", knightProbability(8, 30, 6, 4))
	fmt.Printf("%v", knightProbability(3, 2, 0, 0))
}
