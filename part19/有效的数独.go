package main

func isValidSudoku(board [][]byte) bool {
	rowHas := [9][9]bool{}
	colHas := [9][9]bool{}
	boxHas := [3][3][9]bool{}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				continue
			}
			x := b - '1'
			if rowHas[i][x] || colHas[j][x] || boxHas[i/3][j/3][x] {
				return false
			}
			rowHas[i][x] = true
			colHas[j][x] = true
			boxHas[i/3][j/3][x] = true
		}
	}

	return true
}

func main() {
	println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}
