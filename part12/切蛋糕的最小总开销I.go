package main

import "math"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	dp := make([][][][]int, m)
	for i := range dp {
		dp[i] = make([][][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([][]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, n)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = math.MaxInt
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][i][j][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l1, l2 := 0, i
			for l2 < m {
				r1, r2 := 0, j
				for r2 < n {
					for mid := l1; mid < l2; mid++ {
						dp[l1][l2][r1][r2] = min(dp[l1][l2][r1][r2], dp[l1][mid][r1][r2]+dp[mid+1][l2][r1][r2]+horizontalCut[mid])
					}
					for mid := r1; mid < r2; mid++ {
						dp[l1][l2][r1][r2] = min(dp[l1][l2][r1][r2], dp[l1][l2][r1][mid]+dp[l1][l2][mid+1][r2]+verticalCut[mid])
					}
					r1++
					r2++
				}
				l1++
				l2++
			}
		}
	}
	return dp[0][m-1][0][n-1]
}

//func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
//	dp := make([][][][]int, m)
//	for i := range dp {
//		dp[i] = make([][][]int, m)
//		for j := range dp[i] {
//			dp[i][j] = make([][]int, n)
//			for k := range dp[i][j] {
//				dp[i][j][k] = make([]int, n)
//				for l := range dp[i][j][k] {
//					dp[i][j][k][l] = math.MaxInt
//				}
//			}
//		}
//	}
//	var dfs func(l1, l2, r1, r2 int) int
//	dfs = func(l1, l2, r1, r2 int) int {
//		if dp[l1][l2][r1][r2] != math.MaxInt {
//			return dp[l1][l2][r1][r2]
//		}
//		for i := l1; i < l2; i++ {
//			dp[l1][l2][r1][r2] = min(dp[l1][l2][r1][r2], dfs(l1, i, r1, r2)+dfs(i+1, l2, r1, r2)+horizontalCut[i])
//		}
//		for j := r1; j < r2; j++ {
//			dp[l1][l2][r1][r2] = min(dp[l1][l2][r1][r2], dfs(l1, l2, r1, j)+dfs(l1, l2, j+1, r2)+verticalCut[j])
//		}
//		if dp[l1][l2][r1][r2] == math.MaxInt {
//			dp[l1][l2][r1][r2] = 0
//		}
//		return dp[l1][l2][r1][r2]
//	}
//	return dfs(0, m-1, 0, n-1)
//}

func main() {
	print(minimumCost(3, 2, []int{1, 3}, []int{5}))
}
