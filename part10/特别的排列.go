package main

import "fmt"

func main() {
	fmt.Printf("%v", specialPerm([]int{2, 3, 6}))
}

func specialPerm(nums []int) (res int) {
	n := len(nums)
	u := 1<<n - 1
	cache := make([][]int, u)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(mask, prevIdx int) (res int) {
		if mask == 0 {
			return 1
		}
		p := &cache[mask][prevIdx]
		if *p != -1 {
			return *p
		}
		prev := nums[prevIdx]
		for j, x := range nums {
			if mask>>j&1 == 1 && (prev%x == 0 || x%prev == 0) {
				res += dfs(mask^1<<j, j)
			}
		}
		*p = res
		return
	}
	for i := range n {
		res += dfs(u^(1<<i), i)
	}
	return res % 1_000_000_007
}
