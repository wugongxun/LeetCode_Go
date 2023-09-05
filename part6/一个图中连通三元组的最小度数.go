package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", minTrioDegree(6, [][]int{
		{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6},
	}))
}

func minTrioDegree(n int, edges [][]int) int {
	g := make([]map[int]bool, n)
	h := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x][y], g[y][x] = true, true
		degree[x]++
		degree[y]++
	}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if degree[x] < degree[y] || (degree[x] == degree[y] && x < y) {
			h[x] = append(h[x], y)
		} else {
			h[y] = append(h[y], x)
		}
	}

	res := math.MaxInt32
	for i := 0; i < n; i++ {
		for _, j := range h[i] {
			for _, k := range h[j] {
				if _, ok := g[i][k]; ok {
					res = min(res, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

//func minTrioDegree(n int, edges [][]int) int {
//	adj := make([][]bool, n)
//	for i := 0; i < n; i++ {
//		adj[i] = make([]bool, n)
//	}
//	degree := make([]int, n)
//	for _, edge := range edges {
//		x, y := edge[0]-1, edge[1]-1
//		adj[x][y], adj[y][x] = true, true
//		degree[x]++
//		degree[y]++
//	}
//	res := math.MaxInt32
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			if adj[i][j] {
//				for k := j + 1; k < n; k++ {
//					if adj[i][k] && adj[j][k] {
//						res = min(res, degree[i]+degree[j]+degree[k]-6)
//					}
//				}
//			}
//		}
//	}
//	if res == math.MaxInt32 {
//		return -1
//	}
//	return res
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
