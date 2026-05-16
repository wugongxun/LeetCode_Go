package main

import (
	"math"
	"slices"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m-k+1)
	for i := range res {
		res[i] = make([]int, n-k+1)
	}
	for i := range res {
		for j := range res[i] {
			var a []int
			for _, row := range grid[i : i+k] {
				a = append(a, row[j:j+k]...)
			}
			slices.Sort(a)
			r := math.MaxInt
			for i := 1; i < len(a); i++ {
				if a[i-1] != a[i] {
					r = min(r, int(math.Abs(float64(a[i-1]-a[i]))))
				}
			}
			if r != math.MaxInt {
				res[i][j] = r
			}
		}
	}
	return res
}
