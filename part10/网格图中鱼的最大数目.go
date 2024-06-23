package main

import "fmt"

func main() {
	//print(findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}))
	dfs(1)
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 {
			return 0
		}
		s := grid[i][j]
		grid[i][j] = 0
		for _, d := range dirs {
			s += dfs(i+d.x, j+d.y)
		}
		return s
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(i int) {
	if i >= 10 {
		return
	}
	fmt.Printf("%d\n", i)
	dfs(i + 1)
}
