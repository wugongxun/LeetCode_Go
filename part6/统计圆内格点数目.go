package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", countLatticePoints([][]int{{2, 2, 1}}))
	fmt.Printf("%v", countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
}

func countLatticePoints(circles [][]int) int {
	sort.Slice(circles, func(i, j int) bool {
		return circles[i][2] > circles[j][2]
	})
	res := 0
	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			for _, c := range circles {
				if (x-c[0])*(x-c[0])+(y-c[1])*(y-c[1]) <= c[2]*c[2] {
					res++
					break
				}
			}
		}
	}
	return res
}
