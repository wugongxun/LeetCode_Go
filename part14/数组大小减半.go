package main

import (
	"sort"
)

func minSetSize(arr []int) int {
	freq := make(map[int]int)
	for _, x := range arr {
		freq[x]++
	}
	var vals []int
	for _, v := range freq {
		vals = append(vals, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))
	s := 0
	for i, c := range vals {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	panic("impossible")
}

func main() {
	print(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
