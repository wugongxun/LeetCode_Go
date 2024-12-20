package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
	"slices"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	n := len(queries)
	slices.SortFunc(rooms, func(a, b []int) int { return b[1] - a[1] })
	queryIds := make([]int, n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		queryIds[i] = i
		res[i] = -1
	}
	slices.SortFunc(queryIds, func(a, b int) int { return queries[b][1] - queries[a][1] })
	roomIds := redblacktree.New[int, any]()
	j := 0
	for _, i := range queryIds {
		preferred, minSize := queries[i][0], queries[i][1]
		for j < len(rooms) && rooms[j][1] >= minSize {
			roomIds.Put(rooms[j][0], struct{}{})
			j++
		}
		minAbs := math.MaxInt
		if node, ok := roomIds.Floor(preferred); ok {
			minAbs = preferred - node.Key
			res[i] = node.Key
		}
		if node, ok := roomIds.Ceiling(preferred); ok && node.Key-preferred < minAbs {
			res[i] = node.Key
		}
	}
	return res
}

func main() {
	closestRoom([][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}})
}
