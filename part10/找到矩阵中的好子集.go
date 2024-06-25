package main

func main() {
	goodSubsetofBinaryMatrix([][]int{{0, 1, 1, 0}, {0, 0, 0, 1}, {1, 1, 1, 1}})
}

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	maskToIdx := make(map[int]int)
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}
	for k1, v1 := range maskToIdx {
		for k2, v2 := range maskToIdx {
			if k1&k2 == 0 {
				return []int{min(v1, v2), max(v1, v2)}
			}
		}
	}
	return []int{}
}
