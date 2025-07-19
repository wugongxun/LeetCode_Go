package main

import "slices"

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int { return a[0] - b[0] })
	s, e := 1, 0
	for _, m := range meetings {
		if m[0] > e {
			days -= e - s + 1
			s = m[0]
		}
		e = max(e, m[1])
	}
	days -= e - s + 1
	return days
}
