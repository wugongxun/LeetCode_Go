package main

import "fmt"

func main() {
	fmt.Printf("%v", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	start, end := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > end {
			if !merged {
				res = append(res, []int{start, end})
				merged = true
			}
			res = append(res, interval)
		} else if interval[1] < start {
			res = append(res, interval)
		} else {
			start = min(start, interval[0])
			end = max(end, interval[1])
		}
	}
	if !merged {
		res = append(res, []int{start, end})
	}
	return res
}

//func insert(intervals [][]int, newInterval []int) (res [][]int) {
//	search(&intervals, newInterval)
//	start, end := intervals[0][0], intervals[0][1]
//	for i := 1; i < len(intervals); i++ {
//		start1, end1 := intervals[i][0], intervals[i][1]
//		if start1 > end {
//			res = append(res, []int{start, end})
//			start, end = start1, end1
//		} else {
//			end = max(end, end1)
//		}
//	}
//	res = append(res, []int{start, end})
//	return res
//}
//
//func search(intervals *[][]int, newInterval []int) {
//	l, r := 0, len(*intervals)
//	for l < r {
//		m := (l + r) / 2
//		if (*intervals)[m][0] <= newInterval[0] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//	*intervals = append((*intervals)[:l], append([][]int{newInterval}, (*intervals)[l:]...)...)
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
