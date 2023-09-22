package main

import "fmt"

func main() {
	//fmt.Printf("%v", checkValidGrid([][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}))
	fmt.Printf("%v", checkValidGrid([][]int{{0, 3, 6}, {5, 8, 1}, {2, 7, 4}}))
}

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	dirs := [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	var dfs func(i, j, cur int) bool
	dfs = func(i, j, cur int) bool {
		if i < 0 || i >= n || j < 0 || j >= n || grid[i][j] != cur {
			return false
		}
		if cur == n*n-1 {
			return true
		}
		for _, dir := range dirs {
			if dfs(i+dir[0], j+dir[1], cur+1) {
				return true
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
