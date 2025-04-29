package main

import (
	"container/heap"
	"sort"
)

func main() {
	print(minOperations([]int{2, 11, 10, 1, 3}, 10))
}

func minOperations(nums []int, k int) int {
	h := &hp{nums}
	heap.Init(h)
	res := 0
	for h.IntSlice[0] < k {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		heap.Push(h, min(x, y)*2+max(x, y))
		heap.Fix(h, 0)
		res++
	}
	return res
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() (v any) {
	a := h.IntSlice
	h.IntSlice, v = a[:len(a)-1], a[len(a)-1]
	return
}
