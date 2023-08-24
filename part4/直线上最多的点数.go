package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", maxPoints([][]int{{9, -25}, {-4, 1}, {-1, 5}, {-7, 7}}))
}

func maxPoints(points [][]int) int {
	n := len(points)
	slope := make(map[float32]int)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x, y := points[j][1]-points[i][1], points[j][0]-points[i][0]
			if x == 0 {
				slope[-math.MaxFloat32]++
			} else if y == 0 {
				slope[math.MaxFloat32]++
			} else {
				slope[float32(x)/float32(y)]++
			}
		}
		for key, count := range slope {
			delete(slope, key)
			if count > res {
				res = count
			}
		}
	}
	return res + 1
}
