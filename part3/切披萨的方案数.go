package main

import "fmt"

func main() {
	fmt.Printf("%v", ways([]string{"A..", "AAA", "..."}, 3))
}

func ways(pizza []string, k int) int {
	const mod = 1e9 + 7
	n, m := len(pizza), len(pizza[0])
	ms := NewMatrixSum(pizza)
	dp := make([][][]int, k)
	for t := range dp {
		dp[t] = make([][]int, n)
		for i := range dp[t] {
			dp[t][i] = make([]int, m)
		}
	}
	for t := 0; t < k; t++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if t == 0 {
					if ms.query(i, j, n, m) > 0 {
						dp[t][i][j] = 1
					}
					continue
				}
				res := 0
				for i2 := i + 1; i2 < n; i2++ {
					if ms.query(i, j, i2, m) > 0 {
						res += dp[t-1][i2][j]
					}
				}
				for j2 := 0; j2 < m; j2++ {
					if ms.query(i, j, n, j2) > 0 {
						res += dp[t-1][i][j2]
					}
				}
				dp[t][i][j] = res % mod
			}
		}
	}
	return dp[k-1][0][0]
}

//func ways(pizza []string, k int) int {
//	const mod = 1e9 + 7
//	ms := NewMatrixSum(pizza)
//	m, n := len(pizza), len(pizza[0])
//	cache := make([][][]int, k)
//	for c := range cache {
//		cache[c] = make([][]int, m)
//		for i := range cache[c] {
//			cache[c][i] = make([]int, n)
//			for j := 0; j < n; j++ {
//				cache[c][i][j] = -1
//			}
//		}
//	}
//	var dfs func(int, int, int) int
//	dfs = func(c, i, j int) (res int) {
//		if c == 0 {
//			if ms.query(i, j, m, n) > 0 {
//				return 1
//			}
//			return 0
//		}
//		p := &cache[c][i][j]
//		if *p != -1 {
//			return *p
//		}
//		for j2 := j + 1; j2 < n; j2++ {
//			if ms.query(i, j, m, j2) > 0 {
//				res += dfs(c-1, i, j2)
//			}
//		}
//		for i2 := i + 1; i2 < m; i2++ {
//			if ms.query(i, j, i2, n) > 0 {
//				res += dfs(c-1, i2, j)
//			}
//		}
//		*p = res % mod
//		return *p
//	}
//	return dfs(k-1, 0, 0)
//}

type MatrixSum [][]int

func NewMatrixSum(matrix []string) MatrixSum {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, n+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + int(x&1)
		}
	}
	return sum
}

// 返回左上角在 (r1,c1) 右下角在 (r2-1,c2-1) 的子矩阵元素和（类似前缀和的左闭右开）
func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	return s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
}
