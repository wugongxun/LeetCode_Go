package main

import "container/heap"

func clearStars(s string) string {
	bytes := []byte(s)
	h := &hp{}
	heap.Init(h)
	for i, ch := range s {
		if ch == '*' {
			mn := heap.Pop(h).(item)
			bytes[mn.idx] = '*'
		} else {
			heap.Push(h, item{i, ch})
		}
	}
	var res []byte
	for _, ch := range bytes {
		if ch != '*' {
			res = append(res, ch)
		}
	}
	return string(res)
}

type item struct {
	idx int
	ch  int32
}
type hp []item

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	if h[i].ch == h[j].ch {
		return h[i].idx > h[j].idx
	}
	return h[i].ch < h[j].ch
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(item)) }
func (h *hp) Pop() (v any) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func main() {
	print(clearStars("kddk**"))
}
