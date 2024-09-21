package main

import "math"

type Pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	var from, to []Pair
	for i, row := range grid {
		for j, x := range row {
			if x > 1 {
				for k := 1; k < x; k++ {
					from = append(from, Pair{i, j})
				}
			}
			if x == 0 {
				to = append(to, Pair{i, j})
			}
		}
	}
	res := math.MaxInt
	permute(from, 0, func() {
		total := 0
		for i, f := range from {
			total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
		}
		res = min(res, total)
	})
	return res
}

func permute(a []Pair, i int, do func()) {
	if i == len(a) {
		do()
		return
	}
	for j := i; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, i+1, do)
		a[i], a[j] = a[j], a[i]
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
