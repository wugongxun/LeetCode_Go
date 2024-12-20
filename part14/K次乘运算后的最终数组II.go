package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%v", getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Printf("%v", getFinalState([]int{100000, 2000}, 2, 1000000))
}

const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n, mx := len(nums), 0
	h := make(hp, n)
	for i, num := range nums {
		mx = max(mx, num)
		h[i] = pair{num, i}
	}
	heap.Init(&h)
	for ; k > 0 && h.Peek().(pair).x < mx; k-- {
		p := h.Pop().(*pair)
		p.x *= multiplier
		heap.Fix(&h, 0)
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].less(h[j])
	})
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * quickMul(multiplier, e) % mod
	}
	return nums
}

func quickMul(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

type pair struct{ x, i int }

func (p pair) less(o pair) bool {
	return p.x < o.x || (p.x == o.x && p.i < o.i)
}

type hp []pair

func (h hp) Peek() any {
	return h[0]
}

func (h hp) Push(x any) {
	h = append(h, x.(pair))
}

func (h hp) Pop() (v any) {
	h, v = h[1:], &h[0]
	return
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].less(h[j])
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
