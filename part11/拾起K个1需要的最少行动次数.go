package main

import "math"

func main() {
	print(minimumMoves([]int{1, 1}, 2, 4))
}

func minimumMoves(nums []int, k int, maxChanges int) int64 {
	var pos []int
	c := 0
	for i, num := range nums {
		if num == 0 {
			continue
		}
		pos = append(pos, i)
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3
			} else {
				c = max(c, 2)
			}
		}
	}
	c = min(c, k)
	if maxChanges >= k-c {
		return int64(max(0, c-1) + 2*(k-c))
	}

	n := len(pos)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + pos[i]
	}

	res := math.MaxInt
	size := k - maxChanges

	for r := size; r <= n; r++ {
		l := r - size
		mid := (r + l) / 2
		s1 := pos[mid]*(mid-l) - (prefix[mid] - prefix[l])
		s2 := prefix[r] - prefix[mid] - pos[mid]*(r-mid)
		res = min(res, s1+s2)
	}
	return int64(res + maxChanges*2)
}
