package main

import "bytes"

func main() {
	println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	mat := make([][]byte, numRows)
	n := numRows*2 - 2
	idx := make([]int, n)
	for i := 0; i < numRows; i++ {
		idx[i] = i
	}
	for i := 0; i < numRows-2; i++ {
		idx[i+numRows] = numRows - 2 - i
	}
	index := 0
	for _, ch := range s {
		mat[idx[index%n]] = append(mat[idx[index%n]], byte(ch))
		index++
	}
	return string(bytes.Join(mat, nil))
}
