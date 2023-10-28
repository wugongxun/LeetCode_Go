package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", maxSatisfaction([]int{4, 3, 2}))
	fmt.Printf("%v", maxSatisfaction([]int{-5, -7, 8, -2, 1, 3, 9, 5, -10, 4, -5, -2, -1, -6}))
}

func maxSatisfaction(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	res, prefix := 0, 0
	for _, s := range satisfaction {
		prefix += s
		if prefix < 0 {
			break
		}
		res += prefix
	}
	return res

	//n := len(satisfaction)
	//sort.Ints(satisfaction)
	//res := 0
	//for i := n - 1; i >= 0; i-- {
	//	t := 0
	//	for j := i; j < n; j++ {
	//		t += (j - i + 1) * satisfaction[j]
	//	}
	//	res = max(res, t)
	//}
	//return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
