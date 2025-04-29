package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("%v", combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 27))
}

func combinationSum2(candidates []int, target int) (res [][]int) {
	n := len(candidates)
	var path []int
	slices.Sort(candidates)
	var dfs func(int, int)
	dfs = func(i, r int) {
		if r == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		for j := i; j < n && candidates[j] <= r; j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			path = append(path, candidates[j])
			dfs(j+1, r-candidates[j])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return
}
