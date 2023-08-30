package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", numFactoredBinaryTrees([]int{2, 4}))
	fmt.Printf("%v", numFactoredBinaryTrees([]int{45, 42, 2, 18, 23, 1170, 12, 41, 40, 9, 47, 24, 33, 28, 10, 32, 29, 17, 46, 11, 759, 37, 6, 26, 21, 49, 31, 14, 19, 8, 13, 7, 27, 22, 3, 36, 34, 38, 39, 30, 43, 15, 4, 16, 35, 25, 20, 44, 5, 48}))
}

const mod = 1e9 + 7

func numFactoredBinaryTrees(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	ids := make(map[int]int, n)
	for i, x := range arr {
		ids[x] = i
	}
	dp := make([]int, n)
	res := 0
	for i, v := range arr {
		dp[i] = 1
		for j, x := range arr[:i] {
			if x*x > v {
				break
			}
			if x*x == v {
				dp[i] += dp[j] * dp[j]
				break
			}
			if v%x == 0 {
				k, exist := ids[v/x]
				if exist {
					dp[i] += dp[j] * dp[k] * 2
				}
			}
		}
		res += dp[i] % mod
	}
	return res % mod
}

//func numFactoredBinaryTrees(arr []int) int {
//	n := len(arr)
//	sort.Ints(arr)
//	ids := make(map[int]int, n)
//	for i, x := range arr {
//		ids[x] = i
//	}
//	cache := make([]int, n)
//	for i := range cache {
//		cache[i] = -1
//	}
//	var dfs func(int) int
//	dfs = func(i int) int {
//		p := &cache[i]
//		if *p != -1 {
//			return *p
//		}
//		res, val := 1, arr[i]
//		for j, x := range arr[:i] {
//			if x*x > val {
//				break
//			}
//			if x*x == val {
//				res += dfs(j) * dfs(j)
//				break
//			}
//			if val%x == 0 {
//				k, exist := ids[val/x]
//				if exist {
//					res += dfs(j) * dfs(k) * 2
//				}
//			}
//		}
//		*p = res % mod
//		return *p
//	}
//	res := 0
//	for i := 0; i < n; i++ {
//		res += dfs(i)
//	}
//	return res % mod
//}
