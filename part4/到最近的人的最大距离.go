package main

import "fmt"

func main() {
	fmt.Printf("%v", maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) (res int) {
	prev := -1
	for i, val := range seats {
		if val == 1 {
			if prev == -1 {
				res = max(i, res)
			} else {
				res = max((i-prev)/2, res)
			}
			prev = i
		}
	}
	res = max(len(seats)-prev-1, res)
	return res
}
