package main

func movesToChessboard(board [][]int) int {
	n := len(board)
	firstRow := board[0]
	firstCol := make([]int, n)
	var rowCount, colCount [2]int
	for i, row := range board {
		rowCount[firstRow[i]]++
		firstCol[i] = row[0]
		colCount[firstCol[i]]++
	}
	// 如果第一行或者第一列的0和1的个数相差大于1，不可能变为棋盘
	if abs(rowCount[0]-rowCount[1]) > 1 || abs(colCount[0]-colCount[1]) > 1 {
		return -1
	}

	// 将所有行和第一行比较，要么全部相同，要么全不相同
	for _, row := range board {
		same := row[0] == firstRow[0]
		for i, x := range row {
			if (firstRow[i] == x) != same {
				return -1
			}
		}
	}
	return minSwap(firstRow, rowCount) + minSwap(firstCol, colCount)
}

// 最小交换次数
func minSwap(arr []int, count [2]int) int {
	n := len(arr)
	// 如果1的个数大于0的个数，说明1要放在偶数位置，0要放在奇数位置
	// 反之，0要放在偶数位置，1要放在奇数位置
	x0 := 0
	if count[1] > count[0] {
		x0 = 1
	}
	diff := 0
	for i, x := range arr {
		// i%2表示当前位置是偶数位还是奇数位
		// x0为0表示偶数位置要放0，奇数位置要放1，为1表示偶数位置要放1，奇数位置要放0
		// i%2 ^ x0求出当前位置放0还是1
		// 最终i%2 ^ x ^ x0求出结果为1表示当前位置放的不对，0则是正确的
		diff += i%2 ^ x0 ^ x
	}
	if n%2 == 1 {
		return diff / 2
	}
	return min(diff/2, (n-diff)/2)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
