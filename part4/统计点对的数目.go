package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", countPairs(4, [][]int{
		{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1},
	}, []int{2, 3}))
}

func countPairs(n int, edges [][]int, queries []int) (res []int) {
	degree := make([]int, n)
	cnt := make(map[int]int)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if x > y {
			x, y = y, x
		}
		degree[x]++
		degree[y]++
		cnt[x*n+y]++
	}
	arr := make([]int, n)
	copy(arr, degree)
	sort.Ints(arr)
	for _, query := range queries {
		total := 0
		for i := 0; i < n; i++ {
			j := sort.SearchInts(arr, query-arr[i]+1)
			total += n - max(i+1, j)
		}
		for val, freq := range cnt {
			x, y := val/n, val%n
			if degree[x]+degree[y] > query && degree[x]+degree[y]-freq <= query {
				total--
			}
		}
		res = append(res, total)
	}
	return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
