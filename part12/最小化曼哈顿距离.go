package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
)

func main() {
	print(minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
}

func minimumDistance(points [][]int) int {
	xs := redblacktree.New[int, int]()
	ys := redblacktree.New[int, int]()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		put(xs, x)
		put(ys, y)
	}
	res := math.MaxInt
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		remove(xs, x)
		remove(ys, y)
		res = min(res, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
		put(xs, x)
		put(ys, y)
	}
	return res
}

func put(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	if c == 1 {
		t.Remove(v)
	} else {
		t.Put(v, c-1)
	}
}
