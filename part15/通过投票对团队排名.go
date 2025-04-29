package main

import (
	"cmp"
	"maps"
	"slices"
)

func rankTeams(votes []string) string {
	n := len(votes[0])
	cnts := map[rune][]int{}
	for _, ch := range votes[0] {
		cnts[ch] = make([]int, n)
	}
	for _, vote := range votes {
		for i, ch := range vote {
			cnts[ch][i]++
		}
	}
	res := slices.SortedFunc(maps.Keys(cnts), func(a, b rune) int {
		return cmp.Or(slices.Compare(cnts[b], cnts[a]), cmp.Compare(a, b))
	})
	return string(res)
}
