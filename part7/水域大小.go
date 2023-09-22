package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", pondSizes([][]int{{0, 2, 1, 0}, {0, 1, 0, 1}, {1, 1, 0, 1}, {0, 1, 0, 1, 1}}))
}

func pondSizes(land [][]int) (res []int) {
	n, m := len(land), len(land[0])
	for i, row := range land {
		for j, v := range row {
			if v == 0 {
				count := 0
				q := [][]int{{i, j}}
				land[i][j] = -1
				for len(q) > 0 {
					x, y := q[0][0], q[0][1]
					q = q[1:]
					count++
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							if dx == 0 && dy == 0 {
								continue
							}
							x1, y1 := x+dx, y+dy
							if x1 >= 0 && x1 < n && y1 >= 0 && y1 < m && land[x1][y1] == 0 {
								land[x1][y1] = -1
								q = append(q, []int{x1, y1})
							}
						}
					}
				}
				res = append(res, count)
			}
		}
	}
	sort.Ints(res)
	return
}
