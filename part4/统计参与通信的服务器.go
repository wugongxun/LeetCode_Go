package main

import "fmt"

func main() {
	fmt.Printf("%v", countServers([][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}))
}

func countServers(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rows, cols := make([]int, n), make([]int, m)
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	res := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
				res++
			}
		}
	}
	return res
}
