package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", avoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Printf("%v", avoidFlood([]int{1, 2, 0, 1, 2}))
}

func avoidFlood(rains []int) []int {
	n := len(rains)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	days := []int{}
	isFull := make(map[int]int)
	for i, rain := range rains {
		if rain == 0 {
			days = append(days, i)
		} else {
			res[i] = -1
			if _, ok := isFull[rain]; ok {
				idx := sort.SearchInts(days, isFull[rain])
				if idx == len(days) {
					return []int{}
				}
				res[days[idx]] = rain
				days = append(days[:idx], days[idx+1:]...)
			}
			isFull[rain] = i
		}
	}
	return res
}
