package main

import "slices"

func maxConsecutive(bottom int, top int, special []int) int {
	res := 0
	slices.Sort(special)
	special = append(append([]int{bottom - 1}, special...), top+1)
	for i := 1; i < len(special); i++ {
		res = max(res, special[i]-special[i-1]-1)
	}
	return res
}

func main() {
	println(maxConsecutive(2, 9, []int{4, 6}))
}
