package main

import "container/heap"

func minimumTime(n int, edges [][]int, disappear []int) []int {
	type edge struct{ to, wt int }
	adj := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		adj[x] = append(adj[x], edge{y, w})
		adj[y] = append(adj[y], edge{x, w})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = -1
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx := p.dis
		x := p.x
		if dx > dis[x] {
			continue
		}
		for _, e := range adj[x] {
			y := e.to
			newDis := e.wt + dx
			if newDis < disappear[y] && (newDis < dis[y] || dis[y] < 0) {
				dis[y] = newDis
				heap.Push(&h, pair{newDis, y})
			}
		}
	}
	return dis
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
