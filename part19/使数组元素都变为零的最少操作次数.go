package main

import "math/bits"

func minOperations(queries [][]int) int64 {
	f := func(n int) int {
		m := bits.Len(uint(n))
		k := (m - 1) / 2 * 2
		res := k<<k>>1 - 1<<k/3
		return res + (m+1)/2*(n+1-1<<k)
	}
	res := 0
	for _, q := range queries {
		res += (f(q[1]) - f(q[0]-1) + 1) / 2
	}
	return int64(res)
}

func main() {
	println(minOperations([][]int{{6, 8}}))
}
