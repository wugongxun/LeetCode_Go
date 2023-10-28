package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", maxArea(5, 4, []int{3, 1}, []int{1}))
	fmt.Printf("%v", maxArea(1000000000, 1000000000, []int{2}, []int{2}))
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	mod := 1_000_000_007
	maxWidth, maxHeight := 0, 0
	left, top := 0, 0
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	for _, cut := range verticalCuts {
		maxWidth = max(maxWidth, cut-left)
		left = cut
	}
	maxWidth = max(maxWidth, w-left)
	for _, cut := range horizontalCuts {
		maxHeight = max(maxHeight, cut-top)
		top = cut
	}
	maxHeight = max(maxHeight, h-top)
	return maxWidth * maxHeight % mod
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
