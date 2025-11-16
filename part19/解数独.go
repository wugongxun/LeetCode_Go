package main

import "fmt"

func solveSudoku(board [][]byte) {
	rowHas := [9][9]bool{}
	colHas := [9][9]bool{}
	boxHas := [3][3][9]bool{}
	var empty [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				empty = append(empty, [2]int{i, j})
				continue
			}
			x := b - '1'
			rowHas[i][x] = true
			colHas[j][x] = true
			boxHas[i/3][j/3][x] = true
		}
	}

	var dfs func(k int) bool
	dfs = func(k int) bool {
		if k >= len(empty) {
			return true
		}
		i, j := empty[k][0], empty[k][1]
		for x := 0; x < 9; x++ {
			if rowHas[i][x] || colHas[j][x] || boxHas[i/3][j/3][x] {
				continue
			}
			rowHas[i][x] = true
			colHas[j][x] = true
			boxHas[i/3][j/3][x] = true
			board[i][j] = byte(x + '1')
			if dfs(k + 1) {
				return true
			}
			rowHas[i][x] = false
			colHas[j][x] = false
			boxHas[i/3][j/3][x] = false
		}
		return false
	}

	dfs(0)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)

	for _, row := range board {
		strRow := [9]string{}
		for i, b := range row {
			strRow[i] = string(b)
		}
		fmt.Printf("%v\n", strRow)
	}
}
