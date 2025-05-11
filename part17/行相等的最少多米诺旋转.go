package main

import "math"

func minDominoRotations(tops []int, bottoms []int) int {
	minRot := func(target int) int {
		topCnt, bottomCnt := 0, 0
		for i, x := range tops {
			y := bottoms[i]
			if x != target && y != target {
				return math.MaxInt
			}
			if x != target {
				topCnt++
			}
			if y != target {
				bottomCnt++
			}
		}
		return min(topCnt, bottomCnt)
	}
	res := min(minRot(tops[0]), minRot(bottoms[0]))
	if res == math.MaxInt {
		return -1
	}
	return res
}
