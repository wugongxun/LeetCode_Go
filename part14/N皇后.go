package main

import "fmt"

func solveNQueens(n int) (res [][]string) {
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	var dfs func(y int)
	dfs = func(y int) {
		if y >= n {
			t := make([]string, n)
			for i := range t {
				t[i] = string(grid[i])
			}
			res = append(res, t)
		} else {
			for x := 0; x < n; x++ {
				isOk := true
				for i := y - 1; i >= 0; i-- {
					if grid[i][x] == 'Q' {
						isOk = false
						break
					}
				}
				for i, j := y-1, x-1; isOk && i >= 0 && j >= 0; {
					if grid[i][j] == 'Q' {
						isOk = false
						break
					}
					i--
					j--
				}
				for i, j := y-1, x+1; isOk && i >= 0 && j < n; {
					if grid[i][j] == 'Q' {
						isOk = false
						break
					}
					i--
					j++
				}
				if isOk {
					grid[y][x] = 'Q'
					dfs(y + 1)
					grid[y][x] = '.'
				}
			}
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Printf("%v", solveNQueens(4))
}
