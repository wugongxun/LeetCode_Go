package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", findReplaceString("abcd", []int{2, 0}, []string{"cd", "a"}, []string{"ffff", "eee"}))
	fmt.Printf("%v", findReplaceString("wqzzcbnwxc", []int{5, 2, 7}, []string{"bn", "zzc", "wxc"}, []string{"t", "lwb", "nee"}))
}

func findReplaceString(s string, indices []int, sources, targets []string) string {
	n := len(indices)
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return indices[ids[i]] > indices[ids[j]]
	})
	for _, id := range ids {
		index := indices[id]
		source := sources[id]
		if index+len(source) < len(s) && s[index:index+len(source)] == source {
			s = s[:index] + targets[id] + s[index+len(source):]
		}
	}
	return s
}
