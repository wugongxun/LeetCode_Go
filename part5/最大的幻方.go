package main

import "fmt"

func main() {
	fmt.Printf("%v", largestMagicSquare([][]int{
		{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4},
	}))
}

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rowSum, colSum := make([][]int, n), make([][]int, m)
	for i := 0; i < n; i++ {
		sum := 0
		rowSum[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			sum += grid[i][j]
			rowSum[i][j+1] = sum
		}
	}

	for i := 0; i < m; i++ {
		sum := 0
		colSum[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			sum += grid[j][i]
			colSum[i][j+1] = sum
		}
	}

	k := min(n, m)
	for k > 1 {
		for i := 0; i+k <= n; i++ {
			for j := 0; j+k <= m; j++ {
				if isOk(grid, rowSum, colSum, i, j, k) {
					return k
				}
			}
		}
		k--
	}
	return k
}

func isOk(grid, rowSum, colSum [][]int, i, j, k int) bool {
	eq := rowSum[i][j+k] - rowSum[i][j]
	sum1, sum2 := 0, 0
	for t := 0; t < k; t++ {
		if eq != rowSum[i+t][j+k]-rowSum[i+t][j] || eq != colSum[j+t][i+k]-colSum[j+t][i] {
			return false
		}
		sum1 += grid[i+t][j+t]
		sum2 += grid[i+t][j+k-t-1]
	}
	return sum1 == eq && sum2 == eq
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
